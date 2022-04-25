package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Jack struct{}

func (j Jack) ParseName() string {
	return "jack"
}

func (j Jack) ShortHelp() string {
	return "jack -- force la connexion a un serveur distant"
}

func (j Jack) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(j.ShortHelp() + "\n")
	b.WriteString("USAGE\n")
	b.WriteString("  jack <ADDRESS>\n")
	b.WriteString("ARGUMENTS\n")
	b.WriteString("  ADDRESS -- l'adresse du serveur sur le Net\n")
	return b.String()
}

func (j Jack) Run(ctx Context, args []string) tea.Msg {
	if len(args) < 1 {
		return ParseErrorMsg{
			fmt.Errorf("ADDRESS : %w", errMissingArgument),
			j.LongHelp(),
		}
	}

	// récupérer les arguments
	address := args[0]

	// récupérer le serveur
	server, err := ctx.Game.FindServer(address)
	if err != nil {
		return ErrorMsg{err}
	}

	co := ctx.Console
	co.Server = server
	co.Privilege = 1

	return JackMsg{co}
}

type JackMsg struct {
	Console
}