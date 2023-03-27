package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Identify struct{}

type IdentifyMsg struct {
	Login    string
	Password string
}

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

func (i Identify) Run(args []string) tea.Msg {
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
	return IdentifyMsg{login, password}
}
