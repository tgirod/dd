package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Help affiche l'aide
type Help struct{}

func (c Help) ParseName() string {
	return "help"
}

func (c Help) ShortHelp() string {
	return "help -- affiche l'aide"
}

func (c Help) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(c.ShortHelp() + "\n")
	b.WriteString("USAGE\n")
	b.WriteString("  connect\n")
	return b.String()
}

func (c Help) Run(ctx Context, args []string) tea.Msg {
	b := strings.Builder{}
	b.WriteString("COMMANDES DISPONIBLES\n\n")
	for _, s := range ctx.Console.Node.Sub {
		b.WriteString("  " + s.ShortHelp() + "\n")
	}
	return HelpMsg{b.String()}
}

type HelpMsg struct {
	Help string
}
