package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Data struct{}

func (d Data) ParseName() string {
	return "data"
}

func (d Data) ShortHelp() string {
	return "data -- effectue une recherche sur un service DATABASE"
}

func (d Data) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(d.ShortHelp() + "\n")
	b.WriteString("USAGE\n")
	b.WriteString("  data <DATABASE> <KEYWORD>\n")
	b.WriteString("ARGUMENTS\n")
	b.WriteString("  DATABASE -- nom du service DATABASE\n")
	b.WriteString("  KEYWORD -- mot-clef à rechercher\n")
	return b.String()
}

func (d Data) Run(ctx Context, args []string) tea.Msg {
	if len(args) < 2 {
		return LogMsg{err: errMissingArgument}
	}

	if !ctx.Console.IsConnected() {
		return LogMsg{err: errNotConnected}
	}

	// chercher un service gate avec ce nom
	db, err := ctx.Server.FindDatabase(args[0])
	if err != nil {
		return LogMsg{err: err}
	}

	// a les privilèges suffisant
	if !ctx.Console.HasAccess(db.Restricted) {
		return LogMsg{err: errLowPrivilege}
	}

	// effectuer la recherche
	entries := db.Search(args[1])
	return DataMsg{entries}

}

type DataMsg struct {
	Entries []Entry
}

func (d DataMsg) View() string {
	b := strings.Builder{}

	for i, e := range d.Entries {
		fmt.Fprintf(&b, "%d  %s\n", i, e.Title)
	}

	return b.String()

}