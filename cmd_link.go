package main

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Link struct{}

type LinkListMsg struct{}

type LinkMsg struct {
	Id int
}

func (l Link) ParseName() string {
	return "link"
}

func (l Link) ShortHelp() string {
	return "affiche les liens disponibles ou suit un lien"
}

func (l Link) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(l.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  link [ID]\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  aucun -- liste les liens disponibles\n")
	b.WriteString("  ID    -- suit le lien ID")
	return b.String()
}

func (l Link) Run(c *Client, args []string) tea.Msg {
	cmd := fmt.Sprintf("link %s", strings.Join(args, " "))

	if len(args) == 0 {
		// demander la liste
		return LinkListMsg{}
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return ResultMsg{
			Cmd:   cmd,
			Error: fmt.Errorf("ID : %w", err),
		}
	}

	return LinkMsg{id}
}
