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
	ctx    Context // contexte a exécuter après saisie
	title  string  // titre de la fenêtre modale
	input  textarea.Model
	help   bt_help.Model
	cancel bool
}

func NewText(ctx Context, title string, name string, cancel bool) *TextModel {
	m := TextModel{
		ctx:    ctx,
		title:  title,
		input:  textarea.New(),
		help:   bt_help.New(),
		cancel: cancel,
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

func (m *TextModel) Cancel() (tea.Model, tea.Cmd) {
	args := m.ctx.Args

	// annuler la commande
	if m.cancel || len(args) == 0 {
		return m, MsgToCmd(CloseModalMsg{})
	}

	// retirer le dernier argument et relancer la commande
	m.ctx.Args = args[0 : len(args)-1]
	cmd := tea.Batch(
		MsgToCmd(CloseModalMsg{}),
		MsgToCmd(m.ctx),
	)
	return m, cmd
}