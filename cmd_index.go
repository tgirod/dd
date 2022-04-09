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
	var err error

	if !ctx.Console.IsConnected() {
		return LogMsg{err: errNotConnected}
	}

	// récupérer tous les services associés
	addr := ctx.Console.Server.Address

	// récupérer les links
	msg.Links, err = ListServices[Link](ctx.Game, addr)
	if err != nil {
		return LogMsg{err: err}
	}

	return msg
}

type IndexMsg struct {
	Links []Link
}

func (i IndexMsg) View() string {
	b := strings.Builder{}

	b.WriteString("LINKS\n\n")
	for _, l := range i.Links {
		fmt.Fprintf(&b, "  %s -- %s\n", l.Service.Name, l.Service.Description)
	}
	return b.String()
}
