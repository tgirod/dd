package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Link struct{}

func (l Link) ParseName() string {
	return "link"
}

func (l Link) ShortHelp() string {
	return "link -- utlise un service GATE pour rejoindre un autre serveur"
}

func (l Link) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(l.ShortHelp() + "\n")
	b.WriteString("USAGE\n")
	b.WriteString("  link <GATE>\n")
	b.WriteString("ARGUMENTS\n")
	b.WriteString("  GATE -- nom du service dans le serveur courant\n")
	return b.String()
}

func (l Link) Run(ctx Context, args []string) tea.Msg {
	if len(args) < 1 {
		return LogMsg{err: errMissingArgument}
	}

	if !ctx.Console.IsConnected() {
		return LogMsg{err: errNotConnected}
	}

	// chercher un service gate avec ce nom
	gate, err := ctx.Server.FindGate(args[0])
	if err != nil {
		return LogMsg{err: err}
	}

	// chercher le serveur correspondant
	server, err := ctx.Game.FindServer(gate.TargetAddress)
	if err != nil {
		return LogMsg{err: errServerNotFound}
	}

	// modifier la console pour représenter la nouvelle connexion
	console := ctx.Console
	console.Server = server
	console.Privilege = gate.Privilege

	// envoyer le message pour mettre à jour la console
	return ConnectMsg{console}
}
