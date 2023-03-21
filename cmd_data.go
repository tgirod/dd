package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var Data = Node{
	Name: "data",
	Help: "effectuer une recherche sur le serveur",
	Sub: []Command{
		DataSearch{},
		DataView{},
	},
}

type DataSearch struct{}

func (d DataSearch) ParseName() string {
	return "search"
}

func (d DataSearch) ShortHelp() string {
	return "effectue une recherche avec des mots clefs"
}

func (d DataSearch) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(d.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  data search <KEYWORD>\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  KEYWORD -- mot-clef à rechercher, minimum 3 caractères")
	return b.String()
}

func (d DataSearch) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{
			Cmd:   "data search " + strings.Join(args, " "),
			Error: errNotConnected}
	}

	if len(args) < 1 {
		return ResultMsg{
			Cmd:    "data search " + strings.Join(args, " "),
			Error:  errMissingArgument,
			Output: d.LongHelp(),
		}
	}

	keyword := args[0]

	if len([]rune(keyword)) < 2 {
		return ResultMsg{
			Cmd:    "data search " + strings.Join(args, " "),
			Error:  fmt.Errorf("%s : %w", keyword, errKeywordTooShort),
			Output: d.LongHelp(),
		}
	}

	entries := c.Server.DataSearch(keyword, c.Login)

	// construire la réponse à afficher
	b := strings.Builder{}
	tw := tw(&b)
	fmt.Fprintf(tw, "ID\tKEYWORDS\tTITLE\t\n")
	for _, e := range entries {
		title := e.Title
		fmt.Fprintf(tw, "%s\t%s\t%s\t\n",
			e.ID,
			strings.Join(e.Keywords, " "),
			title,
		)
	}
	tw.Flush()

	return ResultMsg{
		Cmd:    "data search " + strings.Join(args, ""),
		Output: b.String(),
	}
}

type DataView struct{}

func (d DataView) ParseName() string {
	return "view"
}

func (d DataView) ShortHelp() string {
	return "affiche le contenu d'une entrée"
}

func (d DataView) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(d.ShortHelp() + "\n")
	b.WriteString("USAGE\n")
	b.WriteString("  data view <ID>\n")
	b.WriteString("ARGUMENTS\n")
	b.WriteString("  ID -- l'identifiant de l'entrée à afficher\n")
	return b.String()
}

func (d DataView) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{Error: errNotConnected}
	}

	if len(args) < 1 {
		return ResultMsg{
			Cmd:    "data view " + strings.Join(args, " "),
			Error:  errMissingArgument,
			Output: d.LongHelp(),
		}
	}

	id := args[0]
	entry, err := c.Server.FindEntry(id, c.Login)
	if err != nil {
		return ResultMsg{
			Cmd:   "data view " + strings.Join(args, " "),
			Error: err,
		}
	}

	// construire la réponse à afficher
	b := strings.Builder{}
	fmt.Fprintf(&b, "TITLE: %s\n", entry.Title)
	fmt.Fprintf(&b, "KEYWORDS: %s\n", strings.Join(entry.Keywords, " "))
	fmt.Fprintf(&b, "-------------------------------------\n")
	fmt.Fprintf(&b, entry.Content)

	return ResultMsg{
		Cmd:    "data view " + strings.Join(args, " "),
		Output: b.String(),
	}
}
