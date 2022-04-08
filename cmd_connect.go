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
	return "connect -- établit la connexion avec un serveur"
}

func (c Connect) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(c.ShortHelp() + "\n")
	b.WriteString("USAGE\n")
	b.WriteString("\tconnect <ADDRESS> <LOGIN> <PASSWORD>\n")
	b.WriteString("ARGUMENTS\n")
	b.WriteString("\t ADDRESS -- l'adresse du serveur sur le Net\n")
	b.WriteString("\t LOGIN -- identifiant de connexion\n")
	b.WriteString("\t PASSWORD -- mot de passe de connexion\n")
	return b.String()
}

func (c Connect) Run(ctx Context, args []string) tea.Msg {
	if len(args) < 3 {
		return LogMsg{
			errMissingArgument,
			c.LongHelp(),
		}
	}

	// récupérer les arguments
	address := args[0]
	login := args[1]
	password := args[2]

	// récupérer le serveur
	server, err := ctx.Game.FindServer(address)
	if err != nil {
		return LogMsg{err: err}
	}

	privilege, err := server.Connect(login, password)
	if err != nil {
		return LogMsg{
			err: fmt.Errorf("connexion impossible : %w", err),
		}
	}

	return ConnectMsg{server, privilege}
}

// ConnectMsg est retourné quand la connexion est une réussite
type ConnectMsg struct {
	Server        // infos sur le serveur
	Privilege int // niveau de privilège acquis
}
