package main

import (
	bt_help "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type TextKeymap struct {
	Validate key.Binding
	Cancel   key.Binding
}

var DefaultTextKeymap = TextKeymap{
	Validate: key.NewBinding(
		key.WithKeys("tab", "valider"),
		key.WithHelp("tab", "valider"),
	),
	Cancel: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "annuler"),
	),
}

func (k TextKeymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Validate, k.Cancel}
}

func (k TextKeymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Validate},
		{k.Cancel},
	}
}

type TextModel struct {
	ctx   Context // contexte a exécuter après saisie
	cmd   Cmd     // commande pour laquelle la fenêtre est ouverte
	name  string  // nom de l'argument
	title string  // titre de la fenêtre modale
	input textarea.Model
	help  bt_help.Model
}

func NewText(ctx Context, cmd Cmd, name string, title string) *TextModel {
	m := TextModel{
		ctx:   ctx,
		cmd:   cmd,
		name:  name,
		title: title,
		input: textarea.New(),
		help:  bt_help.New(),
	}

	m.input.Placeholder = name
	m.input.SetWidth(UI_WIDTH)
	m.input.SetHeight(20)

	m.help.Width = UI_WIDTH

	return &m
}

func (m *TextModel) Init() tea.Cmd {
	return m.input.Focus()
}

func (m *TextModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultTextKeymap.Validate):
			return m.Validate()
		case key.Matches(msg, DefaultTextKeymap.Cancel):
			return m.Cancel()
		}
		m.input, cmd = m.input.Update(msg)
	}

	return m, cmd
}

func (m *TextModel) View() string {
	title := uiStyle.Copy().Align(lg.Center).MarginBottom(1).Render(m.title)
	input := m.input.View()
	help := uiStyle.Copy().MarginTop(1).Render(m.help.View(DefaultTextKeymap))
	return uiStyle.Render(lg.JoinVertical(lg.Left, title, input, help))
}

// Validate ajoute la saisie au contexte et relance l'exécution
func (m *TextModel) Validate() (tea.Model, tea.Cmd) {
	// stocker dans le contexte
	ctx := m.ctx.WithContext(m.name, m.input.Value(), m.cmd)

	// retourner le contexte à relancer
	cmd := tea.Batch(
		MsgToCmd(CloseModalMsg{}),
		MsgToCmd(ctx),
	)
	return m, cmd
}

func (m *TextModel) Cancel() (tea.Model, tea.Cmd) {
	ctx := m.ctx.Cancel()
	return m, tea.Batch(
		MsgToCmd(CloseModalMsg{}),
		MsgToCmd(ctx),
	)
}