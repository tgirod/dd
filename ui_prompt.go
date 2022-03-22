package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Prompt struct {
	input string
	root  cmd
}

type cmd struct {
	sub map[string]cmd
}

func (p Prompt) Init() tea.Cmd {
	return nil
}

func (p Prompt) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	fmt.Println("prompt update", msg)
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "p" {
			return p, p.runPub
		}
	}
	return p, nil
}

func (p Prompt) View() string {
	return "> " + p.input
}

func (p Prompt) runPub() tea.Msg {
	return Pub{"pub"}
}
