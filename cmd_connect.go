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
	b.WriteString("  connect <ADDRESS> <LOGIN> <PASSWORD>\n")
	b.WriteString("ARGUMENTS\n")
	b.WriteString("   ADDRESS -- l'adresse du serveur sur le Net\n")
	b.WriteString("   LOGIN -- identifiant de connexion\n")
	b.WriteString("   PASSWORD -- mot de passe de connexion\n")
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

	// vérifier l'existence du login
	privilege, err := server.CheckCredentials(login, password)
	if err != nil {
		return LogMsg{err: err}
	}

	// mettre à jour la console
	console := ctx.Console
	console.ServerID = server.ID
	console.Privilege = privilege
	if err := ctx.Game.Update(&console); err != nil {
		fmt.Println(err)
		return LogMsg{err: errInternalError}
	}

	// retourner la console mise à jour pour que le client l'actualise
	return ConnectMsg{console}
}

type ConnectMsg struct {
	Console
}