package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/list"
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

type ArgType int

const (
	Login     ArgType = iota // identifiant utilisateur
	Password                 // mot de passe utilisateur
	Text                     // ligne de texte libre
	LongText                 // texte plus long
	Amount                   // montant (nombre entier)
	MessageId                // identifiant du message
	LinkId                   // identifiant du lien
)

// Arg décrit un argument. Il n'y a pas d'arguments optionnels
type Arg struct {
	Type      ArgType
	Name      string
	ShortHelp string
}

// Context décrit le contexte d'exécution d'une commande
type Context struct {
	*Console
	Path []string
	Args []string
	Cmd
}

// Prompt reconstruit la commande entrée à l'origine
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

func (c Context) Parse() any {
	return c.Cmd.Parse(c)
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

	if len(c.SubCmds) > 0 {
		// trouver la sous-commande à exécuter

		// aucune sous-commande saisie
		if len(ctx.Args) == 0 {
			return Result{
				Prompt: ctx.Prompt(),
				Error:  errMissingCommand,
				Output: c.Help(ctx.Args),
			}
		}

		// chercher les sous-commandes avec le préfixe
		cmds := c.Match(ctx.Args[0])

		// aucune commande ne correspond a ce préfixe
		if len(cmds) == 0 {
			return Result{
				Prompt: ctx.Prompt(),
				Error:  fmt.Errorf("%s : %w", ctx.Args[0], errInvalidCommand),
				Output: c.Help(ctx.Args),
			}
		}

		// sélectionner la première commande qui correspond et poursuivre l'exécution
		ctx.Cmd = cmds[0]
		ctx.Path = append(ctx.Path, ctx.Args[0])
		ctx.Args = ctx.Args[1:]
		return ctx.Parse()
	}

	// on est arrivé au bout de l'arbre, la commande doit être exécutée

	if c.Run == nil {
		// ne devrait pas arriver
		return Result{
			Prompt: ctx.Prompt(),
			Error:  errInternalError,
		}
	}

	// vérifier si tous les arguments sont fournis. Si ce n'est pas le cas, ouvrir une fenêtre modale pour la saisie des arguments manquants
	if len(ctx.Args) < len(c.Args) {
		// trouver le premier argument manquant
		arg := c.Args[len(ctx.Args)]
		// afficher une interface de saisie pour cet argument
		switch arg.Type {
		case Login, Text, Amount:
			mod := NewLine(ctx, arg.ShortHelp, arg.Name, false)
			return OpenModalMsg(mod)
		case Password:
			mod := NewLine(ctx, arg.ShortHelp, arg.Name, true)
			return OpenModalMsg(mod)
		case LongText:
			mod := NewText(ctx, arg.ShortHelp, arg.Name)
			return OpenModalMsg(mod)
		case MessageId:
			messages := ctx.Identity.Messages
			items := make([]list.Item, len(messages))
			for i, m := range messages {
				items[i] = m
			}
			mod := NewList(ctx, items)
			return OpenModalMsg(mod)
		case LinkId:
			links := ctx.Console.Server.Links
			items := make([]list.Item, len(links))
			for i, l := range links {
				items[i] = l
			}
			mod := NewList(ctx, items)
			return OpenModalMsg(mod)
		}
	}

	// lancer l'exécution de la commande et retourner le résultat
	return c.Run(ctx)
}

func (c Context) Resume() any {
	return c.Parse()
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
