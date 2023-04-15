package main

import (
	"fmt"
	"strings"
)

// Cmd est une commande intermédiaire ou terminale dans le prompt
// SubCmds == commande intermédiaire
// Parse == commande terminale
type Cmd struct {
	Path       []string // chemin qui mène à la commande
	Name       string   // nom de la commande
	ShortHelp  string   // phrase d'aide
	SubCmds    []Cmd    // sous-commandes (optionnel)
	Args       []Arg    // arguments (optionnel)
	Connected  bool     // la commande nécessite d'être connecté
	Identified bool     // la commande nécessite d'avoir une identité active
	Run        RunFunc  // fonction exécutée (optionnel)
}

type RunFunc func(ctx Context) any

// Arg décrit un argument. Il n'y a pas d'arguments optionnels
type Arg struct {
	Name      string
	ShortHelp string
}

type Context struct {
	*Console
	Path []string
	Args []string
	Cmd
}

func (c Context) Prompt() string {
	return fmt.Sprintf(
		"%s %s",
		strings.Join(c.Path, " "),
		strings.Join(c.Args, " "),
	)
}

// Result créé un objet Result de base par défaut
func (c Context) Result() Result {
	return Result{
		Prompt: c.Prompt(),
	}
}

// Run exécute la commande dans son contexte
func (c Context) Run() any {
	return c.Cmd.Run(c)
}

// Usage décrit l'utilisation d'une commande
func (c Cmd) Usage() string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "%s %s", strings.ToUpper(strings.Join(c.Path, " ")), c.Name)
	if len(c.SubCmds) > 0 {
		fmt.Fprintf(&b, " <SUBCOMMAND>")
		return b.String()
	}
	for _, arg := range c.Args {
		fmt.Fprintf(&b, " <%s>", strings.ToUpper(arg.Name))
	}
	return b.String()
}

// Match trouve la liste des sous-commandes ayant un préfixe donné
func (c Cmd) Match(prefix string) []Cmd {
	cmds := make([]Cmd, 0, len(c.SubCmds))
	for _, cmd := range c.SubCmds {
		if strings.HasPrefix(cmd.Name, prefix) {
			cmds = append(cmds, cmd)
		}
	}
	return cmds
}

func (c Cmd) FullCmd(args []string) string {
	return fmt.Sprintf("%s %s %s",
		strings.Join(c.Path, " "),
		c.Name,
		strings.Join(args, " "),
	)
}

// CheckArgs vérifie que la commande reçoit le bon nombre d'arguments
func (c Cmd) CheckArgs(args []string) error {
	if len(args) < len(c.Args) {
		return fmt.Errorf("%s : %w",
			c.Args[len(args)].Name,
			errMissingArgument,
		)
	}
	return nil
}

func (c Cmd) Parse(ctx Context) any {
	if ctx.Console.Server == nil && c.Connected {
		return Result{
			Prompt: ctx.Prompt(),
			Error:  errNotConnected,
			Output: c.Help(ctx.Args),
		}
	}

	if ctx.Console.Identity == nil && c.Identified {
		return Result{
			Prompt: ctx.Prompt(),
			Error:  errNotIdentified,
			Output: c.Help(ctx.Args),
		}
	}

	if len(c.SubCmds) == 0 {
		if c.Run == nil {
			// ne devrait pas arriver
			return Result{
				Prompt: ctx.Prompt(),
				Error:  errInternalError,
			}
		}
		// vérifier qu'il y a assez d'arguments
		if err := c.CheckArgs(ctx.Args); err != nil {
			return Result{
				Prompt: ctx.Prompt(),
				Error:  err,
				Output: c.Help(ctx.Args),
			}
		}
		// parser les arguments et retourner un message
		ctx.Cmd = c
		return ctx.Run()
	}

	if len(ctx.Args) == 0 {
		return Result{
			Prompt: ctx.Prompt(),
			Error:  errMissingCommand,
			Output: c.Help(ctx.Args),
		}
	}

	cmds := c.Match(ctx.Args[0])

	if len(cmds) == 0 {
		// aucune commande ne correspond a ce préfixe
		return Result{
			Prompt: ctx.Prompt(),
			Error:  fmt.Errorf("%s : %w", ctx.Args[0], errInvalidCommand),
			Output: c.Help(ctx.Args),
		}
	}

	// continuer l'exécution sur la première commande qui match
	ctx.Path = append(ctx.Path, ctx.Args[0])
	ctx.Args = ctx.Args[1:]
	return cmds[0].Parse(ctx)
}

func (c Cmd) Help(args []string) string {
	if len(args) > 0 && len(c.SubCmds) > 0 {
		cmds := c.Match(args[0])
		if len(cmds) == 0 {
			return errInvalidCommand.Error()
		}
		return cmds[0].Help(args[1:])
	}

	b := strings.Builder{}
	tw := tw(&b)

	if c.Name == "" {
		// cas particulier de la commande racine
		fmt.Fprintf(&b, "COMMANDES\n")
		for _, sub := range c.SubCmds {
			fmt.Fprintf(tw, "\t%s\t%s\t\n", strings.ToUpper(sub.Name), sub.ShortHelp)
		}
		tw.Flush()
		return b.String()
	}

	fmt.Fprintf(tw, "%s : %s\n\n", c.Name, c.ShortHelp)
	fmt.Fprintf(tw, "\t%s\n\n", c.Usage())

	if len(c.SubCmds) > 0 {
		fmt.Fprintf(&b, "SOUS-COMMANDES\n")
		for _, sub := range c.SubCmds {
			fmt.Fprintf(tw, "\t%s\t%s\t\n", strings.ToUpper(sub.Name), sub.ShortHelp)
		}
	}

	if len(c.Args) > 0 {
		fmt.Fprintf(&b, "ARGUMENTS\n")
		for _, arg := range c.Args {
			fmt.Fprintf(tw, "\t<%s>\t%s\t\n", strings.ToUpper(arg.Name), arg.ShortHelp)
		}
	}

	tw.Flush()

	return b.String()
}
