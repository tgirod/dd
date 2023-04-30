package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Cmd est une commande intermédiaire ou terminale dans le prompt
// SubCmds == commande intermédiaire
// Parse == commande terminale
type Cmd struct {
	Name       string  // nom de la commande
	ShortHelp  string  // phrase d'aide
	Args       []Arg   // arguments (optionnel)
	SubCmds    []Cmd   // sous-commandes (optionnel)
	Connected  bool    // la commande nécessite d'être connecté
	Identified bool    // la commande nécessite d'avoir une identité active
	Run        RunFunc // fonction exécutée (optionnel)
}

type RunFunc func(ctx Context) any

// ArgType type d'argument possible
type ArgType int

const (
	ShortArg ArgType = iota
	HiddenArg
	LongArg
	NumberArg
	SelectArg
	SelectNumberArg
)

// Arg décrit un argument. Il n'y a pas d'arguments optionnels
type Arg struct {
	Type      ArgType
	Name      string
	ShortHelp string
	Options   func(ctx Context) []Option
}

type Option struct {
	Desc   string
	Value  any
	Filter string
}

func (a Arg) Parse(arg string) (any, error) {
	switch a.Type {
	case ShortArg, HiddenArg, LongArg, SelectArg:
		return arg, nil
	case NumberArg, SelectNumberArg:
		return strconv.Atoi(arg)
	}
	return arg, nil
}

func (a Arg) OpenModal(ctx Context, cmd Cmd) any {
	switch a.Type {
	case ShortArg:
		return OpenModalMsg(NewLine(ctx, cmd, a.Name, a.ShortHelp, false))
	case HiddenArg:
		return OpenModalMsg(NewLine(ctx, cmd, a.Name, a.ShortHelp, true))
	case LongArg:
		return OpenModalMsg(NewText(ctx, cmd, a.Name, a.ShortHelp))
	case NumberArg:
		return OpenModalMsg(NewNumber(ctx, cmd, a.Name, a.ShortHelp))
	case SelectArg, SelectNumberArg:
		return OpenModalMsg(NewSelect(ctx, cmd, a.Name, a.ShortHelp, a.Options(ctx)))
	}
	return nil
}

// Context décrit le contexte d'exécution d'une commande
type Context struct {
	parent *Context
	key    any
	value  any
	cmd    Cmd
}

// New retourne un nouveau contexte encapsulant le contexte parent
func (c Context) New(key, value any, cmd Cmd) Context {
	return Context{
		parent: &c,
		key:    key,
		value:  value,
		cmd:    cmd,
	}
}

// Value accède à une valeur stockée dans le contexte
func (c Context) Value(key any) any {
	if c.key == key {
		return c.value
	}

	if c.parent != nil {
		return c.parent.Value(key)
	}

	return nil
}

// Prompt reconstruit la commande entrée à l'origine
func (c Context) Prompt() string {
	if c.value == "root" || c.parent == nil {
		return ""
	}

	return fmt.Sprintf("%s %s", c.parent.Prompt(), c.value)
}

// Result créé un objet Result de base par défaut
func (c Context) Result() Result {
	return Result{
		Prompt: c.Prompt(),
	}
}

func (c Context) Cancel() *Context {
	return c.parent
}

func (c Cmd) Parse(ctx Context, args []string) any {
	res := ctx.Result()
	// récupérer la console
	console, _ := ctx.Value("console").(*Console)

	if console.Server == nil && c.Connected {
		res.Error = errNotConnected
		res.Output = c.Help(args)
		return res
	}

	if console.Identity == nil && c.Identified {
		return Result{
			Prompt: ctx.Prompt(),
			Error:  errNotIdentified,
			Output: c.Help(args),
		}
	}

	// parser les arguments
	for _, arg := range c.Args {
		if ctx.Value(arg.Name) != nil {
			// déjà dans le contexte, on passe au suivant
			continue
		}

		if len(args) > 0 {
			// parser un argument
			key := arg.Name
			value, err := arg.Parse(args[0])
			if err != nil {
				return Result{
					Prompt: ctx.Prompt(),
					Error:  errInvalidArgument,
					Output: c.Help(args),
				}
			}
			// retiré l'argument parsé et enregistrer dans le contexte
			args = args[1:]
			ctx = ctx.New(key, value, c)
		} else {
			// ouvrir une fenêtre de saisie
			// BUG si aucun arg n'a été parsé jusqu'ici, ctx contient la mauvaise commande !
			return arg.OpenModal(ctx, c)
		}
	}

	// poursuivre l'exécution avec une sous-commande
	if len(c.SubCmds) > 0 {
		// aucune sous-commande saisie
		// FIXME ouvrir une fenêtre modale qui liste les sous-commandes
		if len(args) == 0 {
			return Result{
				Prompt: ctx.Prompt(),
				Error:  errMissingCommand,
				Output: c.Help(args),
			}
		}

		// chercher les sous-commandes avec le préfixe
		cmds := c.Match(args[0])

		// aucune commande ne correspond a ce préfixe
		if len(cmds) == 0 {
			return Result{
				Prompt: ctx.Prompt(),
				Error:  fmt.Errorf("%s : %w", args[0], errInvalidCommand),
				Output: c.Help(args),
			}
		}

		// sauver le choix parsé
		nextCmd := cmds[0]
		ctx = ctx.New(c.Name, nextCmd.Name, c)

		// poursuivre l'exécution
		return nextCmd.Parse(ctx, args[1:])
	}

	// pas de sous-commande, donc commande terminale
	if c.Run == nil {
		// ne devrait pas arriver
		return Result{
			Prompt: ctx.Prompt(),
			Error:  errInternalError,
		}
	}

	// lancer l'exécution de la commande et retourner le résultat
	return c.Run(ctx)
}

func (c Context) Resume() any {
	return c.cmd.Parse(c, []string{})
}

// Usage décrit l'utilisation d'une commande
func (c Cmd) Usage() string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "%s", c.Name)
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
