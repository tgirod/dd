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
	cmd := fmt.Sprintf("identify %s", strings.Join(args, " "))
	if len(args) < 1 {
		return ResultMsg{
			Error:  fmt.Errorf("LOGIN : %w", errMissingArgument),
			Cmd:    cmd,
			Output: i.LongHelp(),
		}
	}

	if len(args) < 2 {
		return ResultMsg{
			Error:  fmt.Errorf("PASSWORD : %w", errMissingArgument),
			Cmd:    cmd,
			Output: i.LongHelp(),
		}
	}

	// récupérer les arguments
	login := args[0]
	password := args[1]

	if err := c.Identify(login, password); err != nil {
		// échec de l'identification
		return ResultMsg{
			Cmd:   cmd,
			Error: fmt.Errorf("identify : %w", err),
		}
	}

	return ResultMsg{
		Cmd:    cmd,
		Output: fmt.Sprintf("identité établie. Bienvenue, %s\n", login),
	}
}
