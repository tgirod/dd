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
	return "plug -- active l'interface neuronale hors connexion"
}

func (p Plug) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(p.ShortHelp() + "\n")
	b.WriteString("USAGE\n")
	b.WriteString("  plug\n")
	return b.String()
}

func (p Plug) Run(c *Client, args []string) tea.Msg {
	if c.Console.IsConnected() {
		return ResultMsg{
			Error:  errConnected,
			Output: p.LongHelp(),
		}
	}

	c.Console.DNI = true

	return ResultMsg{
		Output: "interface neuronale directe activée",
	}
}