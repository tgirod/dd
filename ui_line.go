package main

import (
	bt_help "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type LineKeymap struct {
	Validate key.Binding
	Cancel   key.Binding
}

var DefaultLineKeymap = LineKeymap{
	Validate: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "valider"),
	),
	Cancel: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "annuler"),
	),
}

func (k LineKeymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Validate, k.Cancel}
}

func (k LineKeymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Validate},
		{k.Cancel},
	}
}

type LineModel struct {
	ctx   Context // contexte a exécuter après saisie
	title string  // titre de la fenêtre modale
	input textinput.Model
	help  bt_help.Model
}

func NewLine(ctx Context, title string, name string, hidden bool) *LineModel {
	m := LineModel{
		ctx:   ctx,
		title: title,
		input: textinput.New(),
		help:  bt_help.New(),
	}

	m.input.Placeholder = name
	m.input.Width = 30

	if hidden {
		m.input.EchoMode = textinput.EchoPassword
	}

	return &m
}

func (m *LineModel) Init() tea.Cmd {
	return m.input.Focus()
}

func (m *LineModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultLineKeymap.Validate):
			return m.Validate()
		case key.Matches(msg, DefaultLineKeymap.Cancel):
			return m.Cancel()
		}
		m.input, cmd = m.input.Update(msg)

	case tea.WindowSizeMsg:
		m.help.Width = msg.Width

	default:
		m.input, cmd = m.input.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m *LineModel) View() string {
	return lg.JoinVertical(lg.Left,
		m.title,
		m.input.View(),
		m.help.View(DefaultLineKeymap),
	)
}

// Validate ajoute la saisie au contexte et relance l'exécution
func (m *LineModel) Validate() (tea.Model, tea.Cmd) {
	// ajouter une argument au contexte
	arg := m.input.Value()
	ctx := m.ctx
	ctx.Args = append(ctx.Args, arg)
	// retourner le contexte à relancer
	cmd := tea.Batch(
		MsgToCmd(CloseModalMsg{}),
		MsgToCmd(ctx),
	)
	return m, cmd
}

func (m *LineModel) Cancel() (tea.Model, tea.Cmd) {
	return m, MsgToCmd(CloseModalMsg{})
}
