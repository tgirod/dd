package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Identify struct{}

func (i Identify) ParseName() string {
	return "identify"
}

func (i Identify) ShortHelp() string {
	return "identify\tvalidation d'identité avec le login/password"
}

func (i Identify) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(i.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  identify <LOGIN> <PASSWORD>\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  LOGIN    -- identifiant utilisateur\n")
	b.WriteString("  PASSWORD -- mot de passe utilisateur\n\n")
	return b.String()
}

func (i Identify) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{
			Error: errNotConnected,
		}
	}

	if len(args) < 1 {
		return ResultMsg{
			fmt.Errorf("LOGIN : %w", errMissingArgument),
			i.LongHelp(),
		}
	}

	if len(args) < 2 {
		return ResultMsg{
			fmt.Errorf("PASSWORD : %w", errMissingArgument),
			i.LongHelp(),
		}
	}

	// récupérer les arguments
	login := args[0]
	password := args[1]
	server := c.Console.Server

	if priv, err := server.CheckCredentials(login, password); err != nil {
		// échec de la connexion
		return ResultMsg{
			Error: fmt.Errorf("identify : %w", err),
		}
	} else {
		// succès de la connexion
		c.Console.Privilege = priv
		c.Console.Login = login
		c.Console.History.Push( Target{c.Console.Server.Address,"",
			priv,login,password} )

		b := strings.Builder{}
		fmt.Fprintf(&b, "identité établie. Bienvenue, %s.\n\n", login)

		return ResultMsg{
			Output: b.String(),
		}
	}
}
