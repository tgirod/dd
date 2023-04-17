package main

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type TextModel struct {
	ctx   Context // contexte a exécuter après saisie
	title string  // titre de la fenêtre modale
	input textarea.Model
}

func NewText(ctx Context, title string, name string) *TextModel {
	m := TextModel{
		ctx:   ctx,
		title: title,
		input: textarea.New(),
	}

	m.input.Placeholder = name
	m.input.SetWidth(40)

	return &m
}

func (m *TextModel) Init() tea.Cmd {
	return m.input.Focus()
}

func (m *TextModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyEnter {
			return m.Validate()
		}
	}
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m *TextModel) View() string {
	return lg.JoinVertical(lg.Left,
		m.title,
		m.input.View(),
	)
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
