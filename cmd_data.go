package main

import (
	"fmt"
	"strings"
	"text/tabwriter"

	tea "github.com/charmbracelet/bubbletea"
)

type DataSearch struct{}

func (d DataSearch) ParseName() string {
	return "search"
}

func (d DataSearch) ShortHelp() string {
	return "search -- effectue une recherche avec des mots clefs"
}

func (d DataSearch) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(d.ShortHelp() + "\n")
	b.WriteString("USAGE\n")
	b.WriteString("  data search <KEYWORD>\n")
	b.WriteString("ARGUMENTS\n")
	b.WriteString("  KEYWORD -- mot-clef à rechercher\n")
	return b.String()
}

func (d DataSearch) Run(ctx Context, args []string) tea.Msg {
	if len(args) < 1 {
		return ParseErrorMsg{
			errMissingArgument,
			d.LongHelp(),
		}
	}

	keyword := args[0]

	if !ctx.Console.IsConnected() {
		return ErrorMsg{errNotConnected}
	}

	// effectuer la recherche
	entries := ctx.Database.Search(keyword, ctx.Console.Privilege)
	return DataSearchMsg{entries}
}

type DataSearchMsg struct {
	Entries []Entry
}

func (d DataSearchMsg) View() string {
	b := strings.Builder{}
	tw := tabwriter.NewWriter(&b, 8, 1, 2, ' ', 0)

	fmt.Fprintf(tw, "ID\tKEYWORDS\tTITLE\t\n")
	for _, e := range d.Entries {
		fmt.Fprintf(tw, "%s\t%s\t%s\t\n",
			e.Key,
			strings.Join(e.Keywords, " "),
			e.Title,
		)
	}
	tw.Flush()

	return b.String()
}
