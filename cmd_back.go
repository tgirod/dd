package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Connect établit la connexion à un serveur
type Back struct{}

func (cmd Back) ParseName() string {
	return "back"
}

func (cmd Back) ShortHelp() string {
	return "quitte le serveur actuel et se reconnecte au serveur précédent"
}

func (cmd Back) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(cmd.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  back")
	return b.String()
}

func (cmd Back) Run(client *Client, args []string) tea.Msg {
	if err := client.Back(); err != nil {
		return ResultMsg{
			Cmd:   "back",
			Error: err,
		}
	}

	return ResultMsg{
		Cmd:    "back",
		Output: fmt.Sprintf("connexion établie à l'adresse %s\n\n", client.Server.Address),
	}
}
