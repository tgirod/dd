package main

import (
	"dd/ui/filler"
	"fmt"
	"strconv"
	"strings"
)

// Cmd est une commande intermédiaire ou terminale dans le prompt
// SubCmds == commande intermédiaire
// Parse == commande terminale
type Cmd struct {
	Path       []string                // chemin qui mène à la commande
	Name       string                  // nom de la commande
	ShortHelp  string                  // phrase d'aide
	SubCmds    []Cmd                   // sous-commandes (optionnel)
	Args       []Arg                   // arguments (optionnel)
	Connected  bool                    // la commande nécessite d'être connecté
	Identified bool                    // la commande nécessite d'avoir une identité active
	Parse      func(args []string) any // fonction exécutée (optionnel)
}

// Arg décrit un argument. Il n'y a pas d'arguments optionnels
type Arg struct {
	Name      string
	ShortHelp string
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

type Context struct {
	Connected  bool
	Identified bool
}

func (c Cmd) Run(ctx Context, args []string) any {
	if c.Connected && !ctx.Connected {
		return Eval{
			Cmd:    c.FullCmd(args),
			Error:  errNotConnected,
			Output: c.Help(args),
		}
	}

	if c.Identified && !ctx.Identified {
		return Eval{
			Cmd:    c.FullCmd(args),
			Error:  errNotIdentified,
			Output: c.Help(args),
		}
	}

	if len(c.SubCmds) == 0 {
		if c.Parse == nil {
			// ne devrait pas arriver
			return Eval{
				Cmd:   c.FullCmd(args),
				Error: errInternalError,
			}
		}
		// vérifier qu'il y a assez d'arguments
		if err := c.CheckArgs(args); err != nil {
			return Eval{
				Cmd:    c.FullCmd(args),
				Error:  err,
				Output: c.Help(args),
			}
		}
		// parser les arguments et retourner un message
		return c.Parse(args)
	}

	if len(args) == 0 {
		return Eval{
			Cmd:    c.FullCmd(args),
			Error:  errMissingCommand,
			Output: c.Help(args),
		}
	}

	cmds := c.Match(args[0])

	if len(cmds) == 0 {
		// aucune commande ne correspond a ce préfixe
		return Eval{
			Cmd:   c.FullCmd(args),
			Error: fmt.Errorf("%s : %w", args[0], errInvalidCommand),
		}
	}

	// continuer l'exécution sur la première commande qui match
	return cmds[0].Run(ctx, args[1:])
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

type QuitMsg struct{}

var quit = Cmd{
	Name:      "quit",
	ShortHelp: "ferme la connexion au serveur courant",
	Connected: true,
	Parse: func(args []string) any {
		return QuitMsg{}
	},
}

type LoadMsg struct {
	Code string
}

var load = Cmd{
	Name:      "load",
	ShortHelp: "charge une nouvelle commande",
	Args: []Arg{
		{
			Name:      "code",
			ShortHelp: "code de la commande",
		},
	},
	Parse: func(args []string) any {
		code := args[0]
		return LoadMsg{code}
	},
}

type JackMsg struct {
	Id int
}

var jack = Cmd{
	Name:      "jack",
	ShortHelp: "force l'accès à un lien",
	Connected: true,
	Args: []Arg{
		{
			Name:      "id",
			ShortHelp: "identifiant du lien",
		},
	},
	Parse: func(args []string) any {
		// récupérer le lien
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return Eval{
				Error: fmt.Errorf("ID : %w", errInvalidArgument),
			}
		}
		return JackMsg{id}
	},
}

type PlugMsg struct{}

var plug = Cmd{
	Name:      "plug",
	ShortHelp: "active l'interface neuronale",
	Parse: func(args []string) any {
		return PlugMsg{}
	},
}

type HelpMsg struct {
	Args []string
}

var help = Cmd{
	Name:      "help",
	ShortHelp: "affiche l'aide",
	Parse: func(args []string) any {
		return HelpMsg{args}
	},
}

type DataSearchMsg struct {
	Keyword string
}

type DataViewMsg struct {
	Id string
}

var data = Cmd{
	Name:      "data",
	ShortHelp: "recherche des données sur le serveur",
	Connected: true,
	SubCmds: []Cmd{
		{
			Name:      "search",
			Path:      []string{"data"},
			ShortHelp: "effectue une recherche par mot clef",
			Args: []Arg{
				{
					Name:      "keyword",
					ShortHelp: "mot clef utilisé pour la recherche",
				},
			},
			Parse: func(args []string) any {
				return DataSearchMsg{args[0]}
			},
		},
		{
			Name:      "view",
			Path:      []string{"data"},
			ShortHelp: "affiche le contenu d'une entrée",
			Args: []Arg{
				{
					Name:      "id",
					ShortHelp: "identifiant de l'entrée à afficher",
				},
			},
			Parse: func(args []string) any {
				return DataViewMsg{args[0]}
			},
		},
	},
}

type LinkListMsg struct{}

type LinkMsg struct {
	Id int
}

var link = Cmd{
	Name:      "link",
	ShortHelp: "utilise les liens pour se connecter à un autre serveur",
	Connected: true,
	SubCmds: []Cmd{
		{
			Name:      "list",
			Path:      []string{"link"},
			ShortHelp: "affiche la liste des liens disponibles",
			Parse: func(args []string) any {
				return LinkListMsg{}
			},
		},
		{
			Name:      "connect",
			Path:      []string{"link"},
			ShortHelp: "suit un lien vers un autre serveur",
			Args: []Arg{
				{
					Name:      "id",
					ShortHelp: "identifiant du lien à suivre",
				},
			},
			Parse: func(args []string) any {
				// récupérer le lien
				id, err := strconv.Atoi(args[0])
				if err != nil {
					return Eval{
						Error: fmt.Errorf("ID : %w", errInvalidArgument),
					}
				}
				return LinkMsg{id}
			},
		},
	},
}

type BackMsg struct{}

var back = Cmd{
	Name:      "back",
	ShortHelp: "quitte le serveur actuel et se reconnecte au serveur précédent",
	Connected: true,
	Parse: func(args []string) any {
		return BackMsg{}
	},
}

type EvadeListMsg struct{}

type EvadeMsg struct {
	Zone string
}

var evade = Cmd{
	Name:      "evade",
	ShortHelp: "effectue une manoeuvre d'évasion pour gagner un peu de temps",
	Connected: true,
	SubCmds: []Cmd{
		{
			Name:      "list",
			Path:      []string{"evade"},
			ShortHelp: "liste les zones mémoires disponibles pour une évasion",
			Parse: func(args []string) any {
				return EvadeListMsg{}
			},
		},
		{
			Name:      "move",
			Path:      []string{"evade"},
			ShortHelp: "effectue la manoeuvre d'evasion vers une zone mémoire",
			Args: []Arg{
				{
					Name:      "zone",
					ShortHelp: "zone mémoire pour l'évasion",
				},
			},
			Parse: func(args []string) any {
				return EvadeMsg{args[0]}
			},
		},
	},
}

type IndexMsg struct{}

var index = Cmd{
	Name:      "index",
	ShortHelp: "liste les services disponibles dans le serveur courant",
	Connected: true,
	Parse: func(args []string) any {
		return IndexMsg{}
	},
}

type ConnectMsg struct {
	Address string
}

var connect = Cmd{
	Name:      "connect",
	ShortHelp: "établit la connexion avec un serveur",
	Args: []Arg{
		{
			Name:      "address",
			ShortHelp: "adresse sur serveur auquel se connecter",
		},
	},
	Parse: func(args []string) any {
		return ConnectMsg{args[0]}
	},
}

type RegistrySearchMsg struct {
	Name string
}

type RegistryEditMsg struct {
	Name string
}

var registry = Cmd{
	Name:      "registry",
	ShortHelp: "liste et manipule les registres du serveur",
	Connected: true,
	SubCmds: []Cmd{
		{
			Name:      "search",
			Path:      []string{"registry"},
			ShortHelp: "recherche dans les registres",
			Args: []Arg{
				{
					Name:      "prefix",
					ShortHelp: "préfixe du nom du registre",
				},
			},
			Parse: func(args []string) any {
				return RegistrySearchMsg{args[0]}
			},
		},
		{
			Name:      "edit",
			Path:      []string{"registry"},
			ShortHelp: "modifie un registre",
			Args: []Arg{
				{
					Name:      "name",
					ShortHelp: "nom du registre à modifier",
				},
			},
			Parse: func(args []string) any {
				return RegistryEditMsg{args[0]}
			},
		},
	},
}

type IdentifyMsg struct {
	Login    string
	Password string
}

func (i IdentifyMsg) SetPassword(password string) filler.PasswordFiller {
	i.Password = password
	return i
}

var identify = Cmd{
	Name:      "identify",
	ShortHelp: "validation d'identité avec le login/password",
	Args: []Arg{
		{
			Name:      "login",
			ShortHelp: "identifiant utilsateur",
		},
	},
	Parse: func(args []string) any {
		msg := IdentifyMsg{Login: args[0]}
		model := filler.New("entrez votre mot de passe", msg)
		return OpenModalMsg(model)
	},
}

type DoorMsg struct{}

var door = Cmd{
	Name:      "door",
	ShortHelp: "créé une backdoor dans le serveur",
	Connected: true,
	Parse: func(args []string) any {
		return DoorMsg{}
	},
}
