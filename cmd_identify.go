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
	return "validation d'identité avec le login/password"
}

func (i Identify) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(i.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  identify <LOGIN> <PASSWORD>\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  LOGIN    -- identifiant utilisateur\n")
	b.WriteString("  PASSWORD -- mot de passe utilisateur")
	return b.String()
}

func (i Identify) Run(c *Client, args []string) tea.Msg {
	if len(args) < 1 {
		return ResultMsg{
			Error:  fmt.Errorf("LOGIN : %w", errMissingArgument),
			Cmd:    "identify " + strings.Join(args, ""),
			Output: i.LongHelp(),
		}
	}

	if len(args) < 2 {
		return ResultMsg{
			Error:  fmt.Errorf("PASSWORD : %w", errMissingArgument),
			Cmd:    "identify " + strings.Join(args, ""),
			Output: i.LongHelp(),
		}
	}

	// récupérer les arguments
	login := args[0]
	password := args[1]

	if err := c.CheckIdentity(login, password); err != nil {
		// échec de l'identification
		return ResultMsg{
			Cmd:   fmt.Sprintf("identify %s %s", login, strings.Repeat("*", len(password))),
			Error: fmt.Errorf("identify : %w", err),
		}
	}

	c.Console.Login = login

	// si on est connecté à un serveur, on tente d'accéder au compte utilisateur
	if c.Console.Server != nil {
		if priv, err := c.CheckAccount(login); err == nil {
			c.Console.Privilege = priv
		}
	}

	b := strings.Builder{}
	fmt.Fprintf(&b, "identité établie. Bienvenue, %s.\n\n", login)

	return ResultMsg{
		Cmd:    fmt.Sprintf("identify %s %s", login, strings.Repeat("*", len(password))),
		Output: b.String(),
	}
}
