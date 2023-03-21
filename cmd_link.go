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
	if !c.Console.IsConnected() {
		return ResultMsg{
			Cmd:   "link " + strings.Join(args, " "),
			Error: errNotConnected,
		}
	}

	if len(args) == 0 {
		// lister les liens disponibles
		b := strings.Builder{}
		tw := tw(&b)
		fmt.Fprintf(tw, "ID\tDESCRIPTION\t\n")
		for i, t := range c.Server.Targets {
			fmt.Fprintf(tw, "%d\t%s\t\n", i, t.Description)
		}
		tw.Flush()

		return ResultMsg{
			Cmd:    "link " + strings.Join(args, " "),
			Output: b.String(),
		}
	}

	// récupérer le lien
	id, err := strconv.Atoi(args[0])
	if err != nil || id < 0 || id >= len(c.Server.Targets) {
		return ResultMsg{
			Cmd:   "link " + strings.Join(args, " "),
			Error: fmt.Errorf("ID : %w", errInvalidArgument),
		}
	}
	target := c.Server.Targets[id]

	// récupérer le serveur correspondant
	server, err := c.Game.FindServer(target.Address)
	if err != nil {
		return ResultMsg{
			Cmd:   "link " + strings.Join(args, " "),
			Error: fmt.Errorf("%s : %w", target.Address, err),
		}
	}

	if priv, err := server.CheckAccount(c.Login); err != nil {
		// échec de la connexion
		return ResultMsg{
			Error: fmt.Errorf("link : %w", err),
			Cmd:   fmt.Sprintf("link %d", id),
		}
	} else {
		// succès de la connexion
		c.Console.Connect(server, priv)
		c.Console.History.Push(Target{server.Address, ""})

		b := strings.Builder{}
		fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", server.Address)
		fmt.Fprintf(&b, "%s\n", server.Description)

		return ResultMsg{
			Cmd:    fmt.Sprintf("link %d", id),
			Output: b.String(),
		}
	}
}
