package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Index struct{}

type IndexMsg struct{}

func (i Index) ParseName() string {
	return "index"
}

func (i Index) ShortHelp() string {
	return "liste les services du serveur courant"
}

func (i Index) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(i.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  index")
	return b.String()
}

func (i Index) Run(c *Client, args []string) tea.Msg {
	return IndexMsg{}
}
