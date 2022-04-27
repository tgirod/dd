package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Rise struct{}

func (r Rise) ParseName() string {
	return "rise"
}

func (r Rise) ShortHelp() string {
	return "rise -- augmente les privilèges d'un niveau"
}

func (r Rise) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(r.ShortHelp() + "\n")
	b.WriteString("USAGE\n")
	b.WriteString("  rise\n")
	return b.String()
}

func (r Rise) Run(ctx Context, args []string) tea.Msg {
	if !ctx.Console.IsConnected() {
		return ErrorMsg{errNotConnected}
	}

	co := ctx.Console
	co.Privilege++
	co.Alarm++
	return RiseMsg{co}
}

type RiseMsg struct {
	Console
}