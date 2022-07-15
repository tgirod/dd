package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
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
	b.WriteString("  connect <ADDRESS>\n")
	b.WriteString("ARGUMENTS\n")
	b.WriteString("  ADDRESS -- l'adresse du serveur sur le Net\n")
	return b.String()
}

func (c Connect) Run(client Client, args []string) tea.Msg {
	if len(args) < 1 {
		return ParseErrorMsg{
			fmt.Errorf("ADDRESS : %w", errMissingArgument),
			c.LongHelp(),
		}
	}

	// récupérer les arguments
	address := args[0]

	// récupérer le serveur
	server, err := client.Game.FindServer(address)
	if err != nil {
		return ErrorMsg{err}
	}

	// construire la fenêtre modale pour demander le login et le password
	modal := ConnectModal{
		Login: Input{
			Focus:       true,
			Placeholder: "login",
			Width:       20,
		},
		Password: Input{
			Placeholder: "password",
			Hidden:      true,
			Width:       20,
		},
		Server: server,
		Client: client,
	}

	return OpenModalMsg(modal)

}

// ConnectModal est l'interface qui demande le login et le password
type ConnectModal struct {
	Login    Input // champ pour saisir le login
	Password Input // champ pour saisir le mot de passe
	Server         // le serveur auquel on tente de se connecter
	Client         // contexte d'exécution de la commande
}

func (c ConnectModal) Init() tea.Cmd {
	return nil
}

func (c ConnectModal) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {

		case tea.KeyTab, tea.KeyShiftTab, tea.KeyUp, tea.KeyDown:
			// inverser le focus
			c.Login.Focus, c.Password.Focus = c.Password.Focus, c.Login.Focus
			return c, nil

		case tea.KeyEnter:
			if c.Login.Focus {
				// passer le focus au champ password
				c.Login.Focus, c.Password.Focus = c.Password.Focus, c.Login.Focus
				return c, nil
			}

			return c, tea.Batch(
				// exécuter la tentative de connexion
				c.Connect,
				// fermer la fenêtre modale
				func() tea.Msg {
					return CloseModalMsg{}
				},
			)

		default:
			// écrire dans le champ qui va bien
			if c.Login.Focus {
				login, cmd := c.Login.Update(msg)
				c.Login = login.(Input)
				return c, cmd
			} else {
				password, cmd := c.Password.Update(msg)
				c.Password = password.(Input)
				return c, cmd
			}
		}
	}

	return c, nil
}

func (c ConnectModal) View() string {
	if c.Login.Focus {
		return lg.JoinVertical(lg.Left,
			focusFieldStyle.Render(c.Login.View()),
			unfocusFieldStyle.Render(c.Password.View()),
		)
	} else {
		return lg.JoinVertical(lg.Left,
			unfocusFieldStyle.Render(c.Login.View()),
			focusFieldStyle.Render(c.Password.View()),
		)
	}
}

func (c ConnectModal) Connect() tea.Msg {
	// vérifier l'existence du login
	privilege, err := c.Server.CheckCredentials(c.Login.Value, c.Password.Value)
	if err != nil {
		return ErrorMsg{err}
	}

	// modifier les infos de la console
	console := c.Console
	console.Server = c.Server
	console.Login = c.Login.Value
	console.Privilege = privilege

	// retourner la console au client
	return ConnectMsg{console}
}

type ConnectMsg struct {
	Console
}
