package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Ad struct {
	Msg string
}

func (a Ad) Init() tea.Cmd {
	fmt.Println("pub init")
	return nil
}

func (a Ad) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return a, a.close
		}
	}
	return a, nil
}

func (a Ad) View() string {
	return "pub: " + a.Msg + "\n" + "press (q) to quit"
}

func (a Ad) close() tea.Msg {
	return Close{}
}
