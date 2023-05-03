package main

import (
	hhelp "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

const UI_WIDTH = 60

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
	cmd   Cmd     // commande pour laquelle la fenêtre est ouverte
	name  string  // utilisé comme clef pour stocker le résultat
	title string  // titre de la fenêtre modale
	input textinput.Model
	help  hhelp.Model
}

func NewLine(ctx Context, cmd Cmd, name string, title string, hidden bool) *LineModel {
	m := LineModel{
		ctx:   ctx,
		cmd:   cmd,
		name:  name,
		title: title,
		input: textinput.New(),
		help:  hhelp.New(),
	}

	m.input.Placeholder = name
	m.input.Width = 30

	if hidden {
		m.input.EchoMode = textinput.EchoPassword
	}

	m.help.Width = UI_WIDTH

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
		default:
			m.input, cmd = m.input.Update(msg)
		}
	default:
		m.input, cmd = m.input.Update(msg)
		return m, cmd
	}

	return m, nil
}

var uiStyle = lg.NewStyle().Width(UI_WIDTH)

func (m *LineModel) View() string {
	title := uiStyle.Copy().Align(lg.Center).MarginBottom(1).Render(m.title)
	input := m.input.View()
	help := uiStyle.Copy().MarginTop(1).Render(m.help.View(DefaultLineKeymap))
	return uiStyle.Render(lg.JoinVertical(lg.Left, title, input, help))
}

// Validate ajoute la saisie au contexte et relance l'exécution
func (m *LineModel) Validate() (tea.Model, tea.Cmd) {
	// stocker dans le contexte
	ctx := m.ctx.WithContext(m.name, m.input.Value(), m.cmd)

	// retourner le contexte à relancer
	cmd := tea.Batch(
		MsgToCmd(CloseModalMsg{}),
		MsgToCmd(ctx),
	)
	return m, cmd
}

func (m *LineModel) Cancel() (tea.Model, tea.Cmd) {
	ctx := m.ctx.Cancel()
	return m, tea.Batch(
		MsgToCmd(CloseModalMsg{}),
		MsgToCmd(ctx),
	)
}
