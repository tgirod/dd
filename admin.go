package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Admin struct {
	prompt textinput.Model
}

func (a Admin) Init() tea.Cmd {
	return textinput.Blink
}

func (a Admin) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return a, tea.Quit
		case tea.KeyEnter:
			value := a.prompt.Value()
			a.prompt.Reset()
			return a, tea.Println(value)
		default:
			a.prompt, cmd = a.prompt.Update(msg)
			return a, cmd
		}
	default:
		a.prompt, cmd = a.prompt.Update(msg)
		return a, cmd
	}
}

func (a Admin) View() string {
	return a.prompt.View()
}

func NewAdmin() Admin {
	prompt := textinput.New()
	prompt.Focus()
	return Admin{
		prompt: prompt,
	}
}

func AdminStart() {
	p := tea.NewProgram(NewAdmin())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
