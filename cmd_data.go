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

type DataSearchMsg struct {
	Keyword string
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
	cmd := fmt.Sprintf("data search %s", strings.Join(args, " "))
	if len(args) < 1 {
		return ResultMsg{
			Cmd:    cmd,
			Error:  fmt.Errorf("KEYWORD : %w", errMissingArgument),
			Output: d.LongHelp(),
		}
	}

	keyword := args[0]
	return DataSearchMsg{keyword}
}

type DataViewMsg struct {
	Id string
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
	cmd := fmt.Sprintf("data view %s", strings.Join(args, " "))

	if len(args) < 1 {
		return ResultMsg{
			Cmd:    cmd,
			Error:  errMissingArgument,
			Output: d.LongHelp(),
		}
	}

	id := args[0]
	return DataViewMsg{id}
}
