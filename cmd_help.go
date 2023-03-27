package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Help affiche l'aide
type Help struct{}

type HelpMsg struct {
	Args []string
}

func (c Help) ParseName() string {
	return "help"
}

func (c Help) ShortHelp() string {
	return "affiche l'aide"
}

func (c Help) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(c.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  help <COMMAND>\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  COMMAND -- nom d'une commande")
	return b.String()
}

func (c Help) Run(client *Client, args []string) tea.Msg {
	return HelpMsg{args}
}
