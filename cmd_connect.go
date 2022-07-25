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
	b.WriteString("  connect <ADDRESS> <LOGIN> <PASSWORD>\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  ADDRESS  -- adresse du serveur sur le Net\n")
	b.WriteString("  LOGIN    -- identifiant utilisateur\n")
	b.WriteString("  PASSWORD -- mot de passe utilisateur")
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

	if len(args) < 2 {
		return ResultMsg{
			Error:  fmt.Errorf("LOGIN : %w", errMissingArgument),
			Cmd:    "connect " + strings.Join(args, " "),
			Output: c.LongHelp(),
		}
	}

	if len(args) < 3 {
		return ResultMsg{
			Error:  fmt.Errorf("PASSWORD : %w", errMissingArgument),
			Cmd:    "connect " + strings.Join(args, " "),
			Output: c.LongHelp(),
		}
	}

	// récupérer les arguments
	address := args[0]
	login := args[1]
	password := args[2]

	// récupérer le serveur
	server, err := client.Game.FindServer(address)
	if err != nil {
		return ResultMsg{
			Error: err,
		}
	}

	if priv, err := server.CheckCredentials(login, password); err != nil {
		// échec de la connexion
		return ResultMsg{
			Error: fmt.Errorf("connect : %w", err),
			Cmd:   fmt.Sprintf("connect %s %s %s", address, login, strings.Repeat("*", len(password))),
		}
	} else {
		// succès de la connexion
		co := client.Console
		co.Privilege = priv
		co.Login = login
		co.Server = server
		co.InitMem()
		co.History.Clear()
		co.History.Push(Target{server.Address, "", priv, login, password})

		b := strings.Builder{}
		fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", server.Address)
		fmt.Fprintf(&b, "%s\n", server.Description)

		return ResultMsg{
			Cmd:    fmt.Sprintf("connect %s %s %s", address, login, strings.Repeat("*", len(password))),
			Output: b.String(),
		}
	}
}
