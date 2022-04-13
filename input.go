package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type Input struct {
	Value       string
	Hidden      bool
	Focus       bool
	Placeholder string
	Width       int
}

func (i Input) Init() tea.Cmd {
	return nil
}

func (i Input) View() string {
	var content string
	value := i.Value
	if i.Hidden {
		l := len([]rune(value))
		value = strings.Repeat("*", l)
	}
	if i.Focus {
		if i.Value == "" {
			content = cursorStyle.Render(i.Placeholder[0:1]) + mutedTextStyle.Render(i.Placeholder[1:])
		} else {
			content = value + cursorStyle.Render(" ")
		}
	} else {
		if i.Value == "" {
			content = mutedTextStyle.Render(i.Placeholder)
		} else {
			content = value
		}
	}

	return lg.PlaceHorizontal(i.Width, lg.Left, content)
}

func (i Input) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// on ne fait rien si on a pas le focus !
	if !i.Focus {
		return i, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyRunes:
			// ajouter la rune au prompt
			i.Value = i.Value + msg.String()
			return i, nil

		case tea.KeyBackspace:
			// supprimer la dernière rune du prompt
			if len(i.Value) > 0 {
				v := i.Value
				i.Value = string([]rune(v)[:len(v)-1])
			}
			return i, nil
		}
	}

	return i, nil
}
