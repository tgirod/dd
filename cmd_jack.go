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

func (j Jack) Run(c *Client, args []string) tea.Msg {
	if len(args) < 1 {
		return ResultMsg{
			fmt.Errorf("ADDRESS : %w", errMissingArgument),
			j.LongHelp(),
		}
	}

	// récupérer les arguments
	address := args[0]

	// récupérer le serveur
	server, err := c.Game.FindServer(address)
	if err != nil {
		return ResultMsg{Error: err}
	}

	co := c.Console
	co.Server = server
	co.Login = "illegal"
	co.Privilege = 1
	co.Alarm++

	return ResultMsg{
		Output: "connexion illégale établie",
	}
}
