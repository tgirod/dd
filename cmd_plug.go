package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Plug struct{}

func (p Plug) ParseName() string {
	return "plug"
}

func (p Plug) ShortHelp() string {
	return "plug\tactive l'interface neuronale hors connexion"
}

func (p Plug) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(p.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  plug\n\n")
	return b.String()
}

func (p Plug) Run(c *Client, args []string) tea.Msg {
	if c.Console.IsConnected() {
		return ResultMsg{
			Cmd:    "plug",
			Error:  errConnected,
			Output: p.LongHelp(),
		}
	}

	c.Console.DNI = true

	return ResultMsg{
		Cmd:    "plug",
		Output: "interface neuronale directe activée",
	}
}
