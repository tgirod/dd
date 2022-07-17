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
	return "search -- effectue une recherche avec des mots clefs"
}

func (d DataSearch) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(d.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  data search <KEYWORD>\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  KEYWORD -- mot-clef à rechercher\n")
	return b.String()
}

func (d DataSearch) Run(c *Client, args []string) tea.Msg {
	if len(args) < 1 {
		return ResultMsg{
			errMissingArgument,
			d.LongHelp(),
		}
	}

	keyword := args[0]

	if !c.Console.IsConnected() {
		return ResultMsg{Error: errNotConnected}
	}

	entries := c.Server.DataSearch(keyword, c.Login)

	// construire la réponse à afficher
	b := strings.Builder{}
	tw := tw(&b)
	fmt.Fprintf(tw, "ID\tKEYWORDS\tTITLE\t\n")
	for _, e := range entries {
		title := e.Title
		if c.Privilege < e.Restricted {
			title = "accès restreint"
		}

		fmt.Fprintf(tw, "%s\t%s\t%s\t\n",
			e.ID,
			strings.Join(e.Keywords, " "),
			title,
		)
	}
	tw.Flush()

	return ResultMsg{
		Output: b.String(),
	}
}

type DataView struct{}

func (d DataView) ParseName() string {
	return "view"
}

func (d DataView) ShortHelp() string {
	return "view -- affiche le contenu d'une entrée"
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
			errMissingArgument,
			d.LongHelp(),
		}
	}

	id := args[0]
	entry, err := c.Server.FindEntry(id, c.Login)
	if err != nil {
		return ResultMsg{
			Error: err,
		}
	}

	if c.Console.Privilege < entry.Restricted {
		return ResultMsg{
			Error: fmt.Errorf("%s : %w", id, errLowPrivilege),
		}
	}

	// construire la réponse à afficher
	b := strings.Builder{}
	fmt.Fprintf(&b, "TITLE: %s\n", entry.Title)
	fmt.Fprintf(&b, "KEYWORDS: %s\n", strings.Join(entry.Keywords, " "))
	fmt.Fprintf(&b, "-------------------------------------\n")
	fmt.Fprintf(&b, entry.Content)

	return ResultMsg{
		Output: b.String(),
	}
}
