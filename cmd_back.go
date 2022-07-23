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

	// récupérer le serveur précédent
	// 1 enlever le lien actuel
	client.Console.History.Pop()
	// 2 have a Peek at the link that allowed to get to previous server
	prev_target, res := client.Console.History.Peek()
	if res != nil {
		// disconnect !!
		client.Console.Server = nil
		client.Console.Login = ""
		client.Console.Privilege = 0
		client.Console.Alert = 0
		client.Console.DNI = false
		client.Console.History.Clear()

		return ResultMsg{
			Cmd:    "back",
			Output: "déconnexion effectuée",
		}

		// return ResultMsg{
		// 	Cmd:   "back",
		// 	Error: res,
		// }
	}

	// récupérer le serveur
	server, err := client.Game.FindServer(prev_target.Address)
	if err != nil {
		return ResultMsg{
			Cmd:   "back",
			Error: err,
		}
	}

	if priv, err := server.CheckCredentials(prev_target.Login, prev_target.Password); err != nil {
		// échec de la connexion
		return ResultMsg{
			Cmd:   "back",
			Error: fmt.Errorf("back : %w", err),
		}
	} else {
		// succès de la connexion
		co := client.Console
		co.Privilege = priv
		co.Login = prev_target.Login
		co.Server = server
		co.Alert = co.Alert / 2 //Alain : back n'est pas sans soucis
		co.InitMem()

		b := strings.Builder{}
		fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", server.Address)
		fmt.Fprintf(&b, "%s\n", server.Description)

		return ResultMsg{
			Cmd:    "back",
			Output: b.String(),
		}
	}
}
