package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// Prompt est le programme principal de la console. Il est lancé
// automatiquement au démarrage et se charge de fournir l'interface permettant
// à l'utilisateur de lancer les autres programmes.
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
			return p, func() tea.Msg { return Ad{"prout de coco channel"} }
		}
	}
	return p, nil
}

func (p Prompt) View() string {
	return "> " + p.input
}
