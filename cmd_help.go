package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Help affiche l'aide
type Help struct{}

func (c Help) ParseName() string {
	return "help"
}

func (c Help) ShortHelp() string {
	return "help -- affiche l'aide"
}

func (c Help) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(c.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  help <COMMAND>\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  COMMAND -- nom d'une commande\n")
	return b.String()
}

func (c Help) Run(client *Client, args []string) tea.Msg {
	if len(args) == 0 {
		b := strings.Builder{}
		b.WriteString("COMMANDES DISPONIBLES\n\n")
		for _, s := range client.Console.Node.Sub {
			b.WriteString("  " + s.ShortHelp() + "\n")
		}
		return ResultMsg{
			Output: b.String(),
		}
	}

	cmd := args[0]
	match := client.Console.Node.Match(cmd)
	if len(match) == 0 {
		return ResultMsg{
			Error: fmt.Errorf("%s : %w", cmd, errInvalidCommand),
		}
	}

	return ResultMsg{
		Output: match[0].LongHelp(),
	}
}
