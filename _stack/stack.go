package stack

import tea "github.com/charmbracelet/bubbletea"

// PushMsg is sent to push a new model to the top of the stack
type PushMsg tea.Model

// PopMsg is sent to remove the model at the top of the stack (if any)
type PopMsg struct{}

// Model contains the stack of models and a default view
type Model struct {
	stack []tea.Model
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case PushMsg:
		m.stack = append(m.stack, msg)
		return m, nil
	case PopMsg:
		if len(m.stack) == 0 {
			return m, nil
		}
		m.stack = m.stack[:len(m.stack)-1]
		return m, nil
	}
	return m, nil
}

func (m Model) View() string {
	if len(m.stack) == 0 {
		return ""
	}
	return m.stack[len(m.stack)-1].View()
}