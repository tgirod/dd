package main

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type LinkList struct{}

func (l LinkList) ParseName() string {
	return "list"
}

func (l LinkList) ShortHelp() string {
	return "list -- liste les liens disponibles"
}

func (l LinkList) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(l.ShortHelp() + "\n")
	b.WriteString("USAGE\n")
	b.WriteString("  link list\n")
	return b.String()
}

func (l LinkList) Run(c Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ErrorMsg{errNotConnected}
	}

	// obtenir la liste des targets
	// TODO filtrer en fonction du niveau "restricted" ?
	msg := LinkListMsg{}
	msg.Targets = c.Gate.Targets

	return msg
}

type LinkListMsg struct {
	Targets []Target
}

func (l LinkListMsg) View() string {
	b := strings.Builder{}
	for i, t := range l.Targets {
		fmt.Fprintf(&b, "%1d %s\n", i, t.Description)
	}
	return b.String()
}

type LinkConnect struct{}

func (l LinkConnect) ParseName() string {
	return "connect"
}

func (l LinkConnect) ShortHelp() string {
	return "connect -- suit un lien pour se connecter à un autre serveur"
}

func (l LinkConnect) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(l.ShortHelp() + "\n")
	b.WriteString("USAGE\n")
	b.WriteString("  link connect <NUM>\n")
	b.WriteString("ARGUMENTS\n")
	b.WriteString("  NUM -- numéro du lien à suivre\n")
	return b.String()
}

func (l LinkConnect) Run(c Client, args []string) tea.Msg {
	if len(args) < 1 {
		return ErrorMsg{errMissingArgument}
	}

	if !c.IsConnected() {
		return ErrorMsg{errNotConnected}
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return ErrorMsg{errInvalidArgument}
	}

	if id < 0 || id > len(c.Targets)-1 {
		return ErrorMsg{errInvalidArgument}
	}

	target := c.Targets[id]

	if c.Privilege < target.Restricted {
		return ErrorMsg{errLowPrivilege}
	}

	// chercher le serveur correspondant
	server, err := c.FindServer(target.Address)
	if err != nil {
		return ErrorMsg{errServerNotFound}
	}

	// modifier la console pour représenter la nouvelle connexion
	co := c.Console
	co.Server = server
	co.Privilege = target.Privilege
	co.Alarm = 1

	// envoyer le message pour mettre à jour la console
	return ConnectMsg{co}
}
