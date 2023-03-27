package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Evade struct{}

type EvadeListMsg struct{}

type EvadeMsg struct {
	Zone string
}

func (e Evade) ParseName() string {
	return "evade"
}

func (e Evade) ShortHelp() string {
	return "manoeuvre d'évasion pour gagner un peu de temps"
}

func (e Evade) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(e.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  evade [ZONE]\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  aucun -- lister les zones mémoire disponibles\n")
	b.WriteString("  ZONE  -- évasion vers la zone mémoire")
	return b.String()
}

func (e Evade) Run(c *Client, args []string) tea.Msg {
	if len(args) == 0 {
		return EvadeListMsg{}
	}

	zone := args[0]
	return EvadeMsg{zone}
}
