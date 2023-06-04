package main

import (
	hhelp "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type LongKeymap struct {
	Validate key.Binding
	Cancel   key.Binding
}

var DefaultLongKeymap = LongKeymap{
	Validate: key.NewBinding(
		key.WithKeys("tab", "valider"),
		key.WithHelp("tab", "valider"),
	),
	Cancel: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "annuler"),
	),
}

func (k LongKeymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Validate, k.Cancel}
}

func (k LongKeymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Validate},
		{k.Cancel},
	}
}

type LongModel struct {
	ctx   Context // contexte a exécuter après saisie
	node  Node    // noeud en cours
	input textarea.Model
	help  hhelp.Model
}

func NewLong(ctx Context, node Node) *LongModel {
	m := LongModel{
		ctx:   ctx,
		node:  node,
		input: textarea.New(),
		help:  hhelp.New(),
	}

	m.input.Placeholder = m.node.String()
	m.input.SetWidth(UI_WIDTH)
	m.input.SetHeight(20)

	m.help.Width = UI_WIDTH

	return &m
}

func (m *LongModel) Init() tea.Cmd {
	return m.input.Focus()
}

func (m *LongModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultLongKeymap.Validate):
			return m.Validate()
		case key.Matches(msg, DefaultLongKeymap.Cancel):
			return m.Cancel()
		}
		m.input, cmd = m.input.Update(msg)
	}

	return m, cmd
}

func (m *LongModel) View() string {
	title := uiStyle.Copy().Align(lg.Center).MarginBottom(1).Render(m.node.Help())
	input := m.input.View()
	help := uiStyle.Copy().MarginTop(1).Render(m.help.View(DefaultLongKeymap))
	return uiStyle.Render(lg.JoinVertical(lg.Left, title, input, help))
}

// Validate ajoute la saisie au contexte et relance l'exécution
func (m *LongModel) Validate() (tea.Model, tea.Cmd) {
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

func (m *LongModel) Cancel() (tea.Model, tea.Cmd) {
	return m, tea.Batch(
		MsgToCmd(CloseModalMsg{}),
	)
}
