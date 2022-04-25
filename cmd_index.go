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
	if !ctx.Console.IsConnected() {
		return ErrorMsg{errNotConnected}
	}

	s := ctx.Console.Server
	b := strings.Builder{}

	fmt.Fprintf(&b,
		"GATE: %s (%d liens)\n",
		s.Gate.Description,
		len(s.Gate.Targets),
	)

	fmt.Fprintf(&b,
		"DATABASE : %s (%d entrées)\n",
		s.Database.Description,
		len(s.Database.Entries),
	)

	return IndexMsg(b.String())
}

type IndexMsg string

func (i IndexMsg) View() string {
	return string(i)
}
