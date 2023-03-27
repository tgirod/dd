package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Quit struct{}

type QuitMsg struct{}

func (q Quit) ParseName() string {
	return "quit"
}

func (q Quit) ShortHelp() string {
	return "ferme la connexion au serveur courant"
}

func (q Quit) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(q.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  quit")
	return b.String()
}

func (q Quit) Run(args []string) tea.Msg {
	return QuitMsg{}
}
