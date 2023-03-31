package main

import (
	"strconv"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type BalanceMsg struct{}

type PayMsg struct {
	To       string
	Amount   int
	Password string
}

var bank = Cmd{
	Name:      "bank",
	ShortHelp: "effectuer des opérations bancaires",
	SubCmds: []Cmd{
		{
			Path:      []string{"bank"},
			Name:      "balance",
			ShortHelp: "affiche le solde du compte",
			Parse: func(args []string) any {
				return BalanceMsg{}
			},
		},
		{
			Path:      []string{"bank"},
			Name:      "pay",
			ShortHelp: "effectue un transfert de monnaie",
			Parse: func(args []string) any {
				return OpenModalMsg(NewPayModel())
			},
		},
	},
}

const (
	PayFieldTo int = iota
	PayFieldAmount
	PayFieldPassword
)

type PayModel struct {
	Fields []textinput.Model
	Keymap struct {
		Next     key.Binding
		Validate key.Binding
		Cancel   key.Binding
	}
}

func NewPayModel() *PayModel {
	p := PayModel{
		Fields: []textinput.Model{
			textinput.New(),
			textinput.New(),
			textinput.New(),
		},
	}

	// to
	p.Fields[PayFieldTo].Width = 20
	p.Fields[PayFieldTo].Placeholder = "destinataire"

	// amount
	p.Fields[PayFieldAmount].Width = 20
	p.Fields[PayFieldAmount].Placeholder = "montant"

	// mot de passe
	p.Fields[PayFieldPassword].Width = 20
	p.Fields[PayFieldPassword].Placeholder = "mot de passe"
	p.Fields[PayFieldPassword].EchoMode = textinput.EchoPassword

	// keymap
	p.Keymap.Next = key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "champ suivant"),
	)
	p.Keymap.Validate = key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "valide la saisie"),
	)
	p.Keymap.Cancel = key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "annule la saisie"),
	)

	return &p
}

func (p *PayModel) Init() tea.Cmd {
	return p.Fields[PayFieldTo].Focus()
}

func (p *PayModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, p.Keymap.Next):
			cmds = append(cmds, p.Next())

		case key.Matches(msg, p.Keymap.Validate):
			if p.Current() != PayFieldPassword {
				cmds = append(cmds, p.Next())
				break
			}

			// destinataire
			to := p.Fields[PayFieldTo].Value()

			// montant
			amount, err := strconv.Atoi(p.Fields[PayFieldAmount].Value())
			if err != nil {
				// montant invalide
				p.Fields[PayFieldAmount].Reset()
				cmds = append(cmds, p.Focus(PayFieldAmount))
				break
			}

			// mot de passe
			password := p.Fields[PayFieldPassword].Value()

			// effectuer le transfert et fermer la fenêtre
			cmds = append(cmds,
				MsgToCmd(CloseModalMsg{}),
				MsgToCmd(PayMsg{to, amount, password}),
			)

		case key.Matches(msg, p.Keymap.Cancel):
			// annuler le transfert et fermer la fenêtre
			cmds = append(cmds, MsgToCmd(CloseModalMsg{}))

		default:
			// transférer le message au champ qui a le focus
			for i := range p.Fields {
				p.Fields[i], cmd = p.Fields[i].Update(msg)
				cmds = append(cmds, cmd)
			}
		}

	default:
		// transférer le message au champ qui a le focus
		for i := range p.Fields {
			p.Fields[i], cmd = p.Fields[i].Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	return p, tea.Batch(cmds...)
}

var payStyle = lg.NewStyle().Width(25)

func (p *PayModel) View() string {
	return payStyle.Render(lg.JoinVertical(lg.Left,
		p.Fields[PayFieldTo].View(),
		p.Fields[PayFieldAmount].View(),
		p.Fields[PayFieldPassword].View(),
	))
}

func (p *PayModel) Current() int {
	for i, field := range p.Fields {
		if field.Focused() {
			return i
		}
	}
	return 0
}

func (p *PayModel) Focus(field int) tea.Cmd {
	var cmds []tea.Cmd
	for i := range p.Fields {
		if i == field {
			cmds = append(cmds, p.Fields[i].Focus())
		} else {
			p.Fields[i].Blur()
		}
	}
	return tea.Batch(cmds...)
}
func (p *PayModel) Next() tea.Cmd {
	next := (p.Current() + 1) % len(p.Fields)
	return p.Focus(next)
}
