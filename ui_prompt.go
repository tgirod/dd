package main

import (
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

// Prompt est le programme principal de la console. Il est lancé
// automatiquement au démarrage et se charge de fournir l'interface permettant
// à l'utilisateur de lancer les autres programmes.
type Prompt struct {
	input string // saisie en cours
	root  cmd    // commande racine
	path  []cmd  // chemin des commandes sélectionnées
}

type cmd struct {
	name  string    // nom de la sous-commande
	sub   []cmd     // les sous-commandes disponibles
	model tea.Model // le modèle à compléter
}

func (p Prompt) Init() tea.Cmd {
	return nil
}

func (p Prompt) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// frappe au clavier
		switch {

		case msg.String() == " " || msg.String() == "tab":
			// tenter d'ajouter une sous-commande au chemin
			cmd := p.cmd()
			match := cmd.match(p.input)

			if len(match) == 0 {
				// erreur, aucune sous-commande ne correspond
				p.input = ""
				return p, nil
			}

			// ajouter la commande au chemin et effacter l'input
			p.path = append(p.path, match[0])
			p.input = ""

			return p, nil

		case msg.Type == tea.KeyRunes:
			// ajouter un caractère normal à la fin
			p.input += msg.String()
			return p, nil

		case msg.Type == tea.KeyBackspace:
			if len(p.input) == 0 {
				if len(p.path) > 0 {
					// dépiler la dernière commande
					p.path = p.path[0 : len(p.path)-1]
					return p, nil
				}
				// prompt vide, on ne fait rien
				return p, nil
			}

			// supprimer la dernière rune
			input := []rune(p.input)
			if len(input) > 0 {
				p.input = string(input[:len(input)-1])
			}
		}
	}
	return p, nil
}

func (p Prompt) View() string {
	b := strings.Builder{}
	for _, p := range p.path {
		b.WriteString(p.name + " ")
	}

	b.WriteString(p.input)
	b.WriteString(cursorStyle.Render(" "))

	cmd := p.cmd()
	if cmd.model == nil {
		// afficher les sous-commandes possibles
		match := cmd.match(p.input)
		for _, m := range match {
			b.WriteString(compStyle.Render(" " + m.name))
		}
		return b.String()
	}

	b.WriteString(" [[arguments]]")
	return b.String()
}

// retourne la dernière commande sélectionnée
func (p Prompt) cmd() cmd {
	cmd := p.root
	if len(p.path) > 0 {
		cmd = p.path[len(p.path)-1]
	}
	return cmd
}

// liste les sous-commandes qui correspondent au préfixe
func (c cmd) match(prefix string) []cmd {
	comp := make([]cmd, 0, len(c.sub))
	for _, sub := range c.sub {
		if strings.HasPrefix(sub.name, prefix) {
			comp = append(comp, sub)
		}
	}
	return comp
}

func NewPrompt() Prompt {
	return Prompt{
		root: cmd{
			sub: []cmd{
				{
					name:  "connect",
					model: Connect{},
				},
				{
					name:  "link",
					model: Link{},
				},
				{
					name: "jack",
					sub: []cmd{
						{
							name:  "connect",
							model: JackConnect{},
						},
						{
							name:  "link",
							model: JackLink{},
						},
					},
				},
			},
		},
	}
}
