package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type BackMsg struct{}

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

func (cmd Back) Run(args []string) tea.Msg {
	return BackMsg{}
}
