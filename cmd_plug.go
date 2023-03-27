package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Plug struct{}

type PlugMsg struct{}

func (p Plug) ParseName() string {
	return "plug"
}

func (p Plug) ShortHelp() string {
	return "active l'interface neuronale hors connexion"
}

func (p Plug) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(p.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  plug")
	return b.String()
}

func (p Plug) Run(c *Client, args []string) tea.Msg {
	return PlugMsg{}
}
