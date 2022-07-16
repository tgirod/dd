package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Index struct{}

func (i Index) ParseName() string {
	return "index"
}

func (i Index) ShortHelp() string {
	return "index -- liste les services du serveur courant"
}

func (i Index) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(i.ShortHelp() + "\n")
	b.WriteString("USAGE\n")
	b.WriteString("  index\n")
	return b.String()
}

func (i Index) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{Error: errNotConnected}
	}

	s := c.Console.Server
	b := strings.Builder{}

	b.WriteString(s.Description)
	b.WriteString("\n\n")
	b.WriteString("LIENS DISPONIBLES\n")
	b.WriteString("DONNEES DISPONIBLES\n")

	return ResultMsg{
		Output: b.String(),
	}
}
