package main

import (
	"errors"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	compStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	errStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	okStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
	cursorStyle = lipgloss.NewStyle().Reverse(true)
	paramStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
)

var (
	errInvalidCommand  = errors.New("commande invalide")
	errMissingCommand  = errors.New("commande manquante")
	errMissingArgument = errors.New("argument manquant")
)

type Console struct {
	network *Network // données partagées du jeu
	width   int      // largeur de l'affichage
	height  int      // hauteur de l'affichage
	input   string   // saisie utilisateur
	log     string   // résultat de la dernière commande
	root    Cmd      // commande racine de la console
}

// commande utilisable par l'utilisateur
type Cmd struct {
	Name string                       // nom de la sous-commande
	Help string                       // description de la commande
	Sub  []Cmd                        // les sous-commandes disponibles
	Args []Arg                        // les arguments requis par la commande
	Run  func(args ...string) tea.Cmd // le code exécuté
}

// argument d'une commande
type Arg struct {
	Name string
	Help string
}

func (c Console) Init() tea.Cmd {
	return nil
}

// LogMsg contient le retour d'un programme à ajouter dans les logs
type LogMsg struct {
	err error
	msg string
}

func (c Console) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c":
			// quitter l'application client
			return c, tea.Quit

		case "enter":
			// lancer l'exécution de la commande
			cmd := c.Parse()
			//c.logs = append(c.logs, "> "+c.input)
			c.input = ""
			return c, cmd
		}

		if msg.Type == tea.KeyRunes {
			// ajouter dans le champ input
			c.input += msg.String()
		}

		if msg.Type == tea.KeyBackspace {
			if len(c.input) == 0 {
				return c, nil
			}

			// supprimer la dernière rune
			input := []rune(c.input)
			if len(input) > 0 {
				c.input = string(input[:len(input)-1])
			}
		}

	case tea.WindowSizeMsg:
		// gère le redimensionnement de la fenêtre
		c.height = msg.Height
		c.width = msg.Width
		return c, nil

	case LogMsg:
		// ajoute dans les logs
		c.log = msg.View()
		//c.logs = append(c.logs, string(msg))
		return c, nil
	}

	return c, nil
}

// View délègue l'affichage à la vue au sommet de la pile
func (c Console) View() string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "%s\n", c.log)
	//for _, l := range c.logs {
	//fmt.Fprintln(&b, l)
	//}
	fmt.Fprintf(&b, "> %s\n", c.input)
	return b.String()
}

func NewConsole(n *Network, width, height int) Console {
	c := Console{
		network: n,
		width:   width,
		height:  height,
	}

	c.root = Cmd{
		Name: ">",
		Help: "accédez au Net en toute sécurité",
		Sub: []Cmd{
			{
				Name: "connect",
				Help: "établit la connexion à un serveur via son adresse",
				Args: []Arg{
					{Name: "address", Help: "adresse du serveur"},
					{Name: "login", Help: "identifiant"},
					{Name: "password", Help: "mot de passe"},
				},
				Run: func(args ...string) tea.Cmd {
					return func() tea.Msg {
						return c.Connect(args[0], args[1], args[2])
					}
				},
			},
		},
	}

	return c
}

// parse interprète la saisie utilisateur et retourne la commande à exécuter à
// partir de cette saisie
func (c Console) Parse() tea.Cmd {
	args := strings.Fields(c.input)

	if len(args) == 0 {
		// afficher l'aide
		return func() tea.Msg {
			return LogMsg{
				err: errMissingCommand,
				msg: c.Usage(),
			}
		}
	}

	return c.root.Parse(args)
}

func (c Cmd) Parse(args []string) tea.Cmd {
	if c.IsLeaf() {
		// il n'y a pas de sous-commandes

		if len(args) < len(c.Args) {
			// il manque des arguments
			return func() tea.Msg {
				err := errMissingArgument
				for i := len(c.Args) - 1; i >= len(args); i-- {
					err = fmt.Errorf("%s : %w", c.Args[i].Name, err)
				}
				return LogMsg{
					err: err,
					msg: c.Usage(),
				}
			}
		}

		// on exécute le code de la fonction Run
		return c.Run(args...)
	}

	// rechercher une sous-commande
	match := c.Match(args[0])
	if len(match) == 0 {

		// aucune commande ne correspond, afficher l'aide de cmd
		return func() tea.Msg {
			return LogMsg{
				err: fmt.Errorf("%s : %w", args[0], errInvalidCommand),
				msg: c.Usage(),
			}
		}
	}

	// poursuivre le parsing dans la sous-commande
	return match[0].Parse(args[1:])
}

func (c Cmd) IsLeaf() bool {
	return len(c.Sub) == 0
}

func (c Cmd) Match(prefix string) []Cmd {
	sub := make([]Cmd, 0, len(c.Sub))
	for _, s := range c.Sub {
		if strings.HasPrefix(s.Name, prefix) {
			sub = append(sub, s)
		}
	}
	return sub
}

func (c Cmd) Usage() string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "%s - %s\n", strings.ToUpper(c.Name), c.Help)
	fmt.Fprintln(&b, "USAGE")

	if c.IsLeaf() {
		// pas de sous-commande, afficher les arguments
		fmt.Fprintf(&b, "\t%s", c.Name)
		for _, a := range c.Args {
			fmt.Fprintf(&b, " <%s>", a.Name)
		}
		fmt.Fprintln(&b, "")
		fmt.Fprintln(&b, "ARGUMENTS")
		for _, a := range c.Args {
			fmt.Fprintf(&b, "\t%s\n", a.Usage())
		}
		return b.String()
	}

	// afficher la liste des sous-commandes
	fmt.Fprintf(&b, "%s <COMMANDE>\n", c.Name)
	fmt.Fprintln(&b, "COMMANDES")
	for _, s := range c.Sub {
		fmt.Fprintf(&b, "%12s - %s\n", s.Name, s.Help)
	}
	return b.String()
}

func (a Arg) Usage() string {
	return fmt.Sprintf("%12s - %s", a.Name, a.Help)
}

func (l LogMsg) View() string {
	b := strings.Builder{}
	if l.err != nil {
		fmt.Fprintf(&b, "ERR : %s\n\n", l.err.Error())
	}
	fmt.Fprintf(&b, "%s\n", l.msg)
	return b.String()
}

func (c Console) Usage() string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "Liste des commandes disponibles\n\n")
	for _, cmd := range c.root.Sub {
		fmt.Fprintf(&b, "%12s - %s\n", cmd.Name, cmd.Help)
	}
	return b.String()
}
