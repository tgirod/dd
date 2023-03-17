package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Connect établit la connexion à un serveur
type Connect struct{}

func (c Connect) ParseName() string {
	return "connect"
}

func (c Connect) ShortHelp() string {
	return "établit la connexion avec un serveur"
}

func (c Connect) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(c.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  connect <ADDRESS>\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  ADDRESS  -- adresse du serveur sur le Net\n")
	return b.String()
}

func (c Connect) Run(client *Client, args []string) tea.Msg {
	if len(args) < 1 {
		return ResultMsg{
			Error:  fmt.Errorf("ADDRESS : %w", errMissingArgument),
			Cmd:    "connect " + strings.Join(args, " "),
			Output: c.LongHelp(),
		}
	}

	// récupérer les arguments
	address := args[0]

	// récupérer le serveur
	server, err := client.Game.FindServer(address)
	if err != nil {
		return ResultMsg{
			Error: err,
		}
	}

	if priv, err := server.CheckAccount(client.Login); err != nil {
		// échec de la connexion
		return ResultMsg{
			Error: fmt.Errorf("connect : %w", err),
			Cmd:   fmt.Sprintf("connect %s", address),
		}
	} else {
		// succès de la connexion
		client.Console.Connect(server, priv)
		client.Console.History.Clear()
		client.Console.History.Push(Target{server.Address, "", priv})

		b := strings.Builder{}
		fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", server.Address)
		fmt.Fprintf(&b, "%s\n", server.Description)

		return ResultMsg{
			Cmd:    fmt.Sprintf("connect %s", address),
			Output: b.String(),
		}
	}
}
