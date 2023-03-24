package main

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Link struct{}

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
		b := strings.Builder{}
		tw := tw(&b)
		fmt.Fprintf(tw, "ID\tDESCRIPTION\t\n")
		for i, t := range c.Server.Targets {
			fmt.Fprintf(tw, "%d\t%s\t\n", i, t.Description)
		}
		tw.Flush()

		return ResultMsg{
			Cmd:    cmd,
			Output: b.String(),
		}
	}

	// récupérer le lien
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return ResultMsg{
			Cmd:   cmd,
			Error: fmt.Errorf("ID : %w", err),
		}
	}

	if err := c.Link(id); err != nil {
		return ResultMsg{
			Cmd:   cmd,
			Error: err,
		}
	}

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", c.Server.Address)
	fmt.Fprintf(&b, "%s\n", c.Server.Description)

	return ResultMsg{
		Cmd:    cmd,
		Output: b.String(),
	}
}
