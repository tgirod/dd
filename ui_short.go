package main

import (
	hhelp "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

const UI_WIDTH = 60

type ShortKeymap struct {
	Validate key.Binding
	Cancel   key.Binding
}

var DefaultShortKeymap = ShortKeymap{
	Validate: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "valider"),
	),
	Cancel: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "annuler"),
	),
}

func (k ShortKeymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Validate, k.Cancel}
}

func (k ShortKeymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Validate},
		{k.Cancel},
	}
}

type ShortModel struct {
	ctx   Context // contexte a exécuter après saisie
	node  Node    // noeud en cours
	input textinput.Model
	help  hhelp.Model
}

func NewShort(ctx Context, node Node, hidden bool) *ShortModel {
	m := ShortModel{
		ctx:   ctx,
		node:  node,
		input: textinput.New(),
		help:  hhelp.New(),
	}

	m.input.Placeholder = node.String()
	m.input.Width = 30

	if hidden {
		m.input.EchoMode = textinput.EchoPassword
	}

	m.help.Width = UI_WIDTH

	return &m
}

func (m *ShortModel) Init() tea.Cmd {
	return m.input.Focus()
}

func (m *ShortModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultShortKeymap.Validate):
			return m.Validate()
		case key.Matches(msg, DefaultShortKeymap.Cancel):
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

func (m *ShortModel) View() string {
	title := uiStyle.Copy().Align(lg.Center).MarginBottom(1).Render(m.node.Help())
	input := m.input.View()
	help := uiStyle.Copy().MarginTop(1).Render(m.help.View(DefaultShortKeymap))
	return uiStyle.Render(lg.JoinVertical(lg.Left, title, input, help))
}

// Validate ajoute la saisie au contexte et relance l'exécution
func (m *ShortModel) Validate() (tea.Model, tea.Cmd) {
	// stocker dans le contexte
	ctx := m.ctx.WithContext(
		m.node,
		m.node.String(),
		m.input.Value(),
	)

	// retourner le contexte à relancer
	cmd := tea.Batch(
		MsgToCmd(CloseModalMsg{}),
		MsgToCmd(ctx),
	)
	return m, cmd
}

func (m *ShortModel) Cancel() (tea.Model, tea.Cmd) {
	return m, tea.Batch(
		MsgToCmd(CloseModalMsg{}),
	)
}
