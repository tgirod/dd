package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Quit struct{}

func (q Quit) ParseName() string {
	return "quit"
}

func (q Quit) ShortHelp() string {
	return "quit -- ferme la connexion au serveur courant"
}

func (q Quit) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(q.ShortHelp() + "\n")
	b.WriteString("USAGE\n")
	b.WriteString("  quit\n")
	return b.String()
}

func (q Quit) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{Error: errNotConnected}
	}

	c.Console.Server = nil
	c.Console.Login = ""
	c.Console.Privilege = 0
	c.Console.Alarm = 0

	return ResultMsg{
		Output: "déconnexion effectuée",
	}
}
