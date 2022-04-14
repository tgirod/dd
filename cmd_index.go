package main

import (
	"fmt"
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

func (i Index) Run(ctx Context, args []string) tea.Msg {
	var msg IndexMsg

	if !ctx.Console.IsConnected() {
		return LogMsg{err: errNotConnected}
	}

	// récupérer les services gate
	msg.Gates = ctx.Console.Server.Gates

	return msg
}

type IndexMsg struct {
	Gates []Gate
}

func (i IndexMsg) View() string {
	b := strings.Builder{}

	b.WriteString("GATES\n\n")
	for _, g := range i.Gates {
		fmt.Fprintf(&b, "  %s -- %s\n", g.Service.Name, g.Service.Description)
	}
	return b.String()
}
