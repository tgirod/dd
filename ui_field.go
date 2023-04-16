package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type FieldModel struct {
	ctx   Context // contexte a exécuter après saisie
	title string  // titre de la fenêtre modale
	input textinput.Model
}

func NewField(ctx Context, title string, name string, hidden bool) *FieldModel {
	m := FieldModel{
		ctx:   ctx,
		title: title,
		input: textinput.New(),
	}

	m.input.Placeholder = name
	m.input.Width = 30

	if hidden {
		m.input.EchoMode = textinput.EchoPassword
	}

	return &m
}

func (m *FieldModel) Init() tea.Cmd {
	return m.input.Focus()
}

func (m *FieldModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m *FieldModel) View() string {
	return lg.JoinVertical(lg.Left,
		m.title,
		m.input.View(),
	)
}

// Validate ajoute la saisie au contexte et relance l'exécution
func (m *FieldModel) Validate() (tea.Model, tea.Cmd) {
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
