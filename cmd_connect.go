package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Connect établit la connexion à un serveur
type Connect struct{}

type ConnectMsg struct {
	Address string
}

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

func (c Connect) Run(args []string) tea.Msg {
	cmd := fmt.Sprintf("connect %s", strings.Join(args, " "))

	if len(args) < 1 {
		return ResultMsg{
			Error:  fmt.Errorf("ADDRESS : %w", errMissingArgument),
			Cmd:    cmd,
			Output: c.LongHelp(),
		}
	}

	// récupérer les arguments
	address := args[0]

	return ConnectMsg{
		Address: address,
	}
}
