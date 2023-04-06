package filler

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type LoginFiller interface {
	SetLogin(string) LoginFiller
}

type PasswordFiller interface {
	SetPassword(string) PasswordFiller
}

type FilledMsg struct{}

type Model struct {
	Title  string            // titre de la fenêtre modale
	Msg    tea.Msg           // message à remplir
	Fields []textinput.Model // champs de saisie
}

type keymap struct {
	Next     key.Binding
	Validate key.Binding
	Cancel   key.Binding
}

var keys = keymap{
	Next: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "champ suivant"),
	),
	Validate: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "valide la commande"),
	),
	Cancel: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "annule la commande"),
	),
}

func New(title string, msg tea.Msg) *Model {
	m := Model{
		Title:  title,
		Msg:    msg,
		Fields: []textinput.Model{},
	}

	if _, ok := msg.(LoginFiller); ok {
		login := textinput.New()
		login.Placeholder = "login"
		m.Fields = append(m.Fields, login)
	}

	if _, ok := msg.(PasswordFiller); ok {
		password := textinput.New()
		password.Placeholder = "password"
		password.EchoMode = textinput.EchoPassword
		m.Fields = append(m.Fields, password)
	}

	return &m
}

func (m *Model) current() int {
	for i, field := range m.Fields {
		if field.Focused() {
			return i
		}
	}
	return 0
}

func (m *Model) next() tea.Cmd {
	var cmds []tea.Cmd
	next := (m.current() + 1) % len(m.Fields)

	for i := range m.Fields {
		if i == next {
			cmds = append(cmds, m.Fields[i].Focus())
		} else {
			m.Fields[i].Blur()
		}
	}

	return tea.Batch(cmds...)
}

func (m *Model) Init() tea.Cmd {
	return m.Fields[0].Focus()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Next):
			cmds = append(cmds, m.next())

		case key.Matches(msg, keys.Validate):
			// passer au champ suivant si on n'est pas sur le dernier
			if m.current() != len(m.Fields)-1 {
				cmds = append(cmds, m.next())
				break
			}

			// récupérer la valeur du champ login si il existe
			if f, ok := m.Msg.(LoginFiller); ok {
				login := m.Fields[0].Value()
				m.Msg = f.SetLogin(login)
			}

			// récupérer la valeur du champ password si il existe
			if f, ok := m.Msg.(PasswordFiller); ok {
				password := m.Fields[0].Value()
				m.Msg = f.SetPassword(password)
			}

			// retourner le message complété pour qu'il soit traité
			// et un message de cloture
			cmds = append(cmds,
				func() tea.Msg { return m.Msg },
				func() tea.Msg { return FilledMsg{} },
			)

		case key.Matches(msg, keys.Cancel):
			// effacer le champ Login si il existe
			if f, ok := m.Msg.(LoginFiller); ok {
				m.Msg = f.SetLogin("")
			}

			// effacer le champ password si il existe
			if f, ok := m.Msg.(PasswordFiller); ok {
				m.Msg = f.SetPassword("")
			}

			// retourner le message
			cmds = append(cmds,
				func() tea.Msg { return m.Msg },
				func() tea.Msg { return FilledMsg{} },
			)

		default:
			// transférer aux champs de saisie
			for i := range m.Fields {
				m.Fields[i], cmd = m.Fields[i].Update(msg)
				cmds = append(cmds, cmd)
			}
		}

	default:
		// transférer aux champs de saisie
		for i := range m.Fields {
			m.Fields[i], cmd = m.Fields[i].Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) View() string {
	lines := []string{m.Title}
	for i := range m.Fields {
		lines = append(lines, m.Fields[i].View())
	}

	return lg.JoinVertical(lg.Left,
		lines...,
	)
}
