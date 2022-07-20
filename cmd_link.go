package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Link struct{}

func (l Link) ParseName() string {
	return "link"
}

func (l Link) ShortHelp() string {
	return "link -- affiche les liens disponibles ou suit un lien"
}

func (l Link) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(l.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  link [ADDRESS]\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  aucun   -- liste les liens disponibles\n")
	b.WriteString("  ADDRESS -- suit le lien ID\n")
	return b.String()
}

func (l Link) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{
			Error: errNotConnected,
		}
	}

	if len(args) == 0 {
		// lister les liens disponibles
		b := strings.Builder{}
		tw := tw(&b)
		fmt.Fprintf(tw, "ADDRESS\tDESCRIPTION\t\n")
		for _, t := range c.Server.Targets {
			if c.Console.Privilege >= t.Restricted {
				fmt.Fprintf(tw, "%s\t%s\t\n", t.Address, t.Description)
			} else {
				fmt.Fprintf(tw, "\t%s\t%s\t\n", t.Address, "Accès restreint")
			}
		}
		tw.Flush()

		return ResultMsg{
			Output: b.String(),
		}
	}

	// récupérer le lien
	address := args[0]
	target, err := c.Server.FindTarget(address)
	if err != nil {
		return ResultMsg{
			Error: fmt.Errorf("%s : %w", address, err),
		}
	}

	// vérifier le niveau de privilège
	if c.Console.Privilege < target.Restricted {
		return ResultMsg{
			Error: errLowPrivilege,
		}
	}

	// récupérer le serveur correspondant
	server, err := c.Game.FindServer(address)
	if err != nil {
		return ResultMsg{
			Error: fmt.Errorf("%s : %w", address, err),
		}
	}

	// effectuer la connexion avec le serveur
	priv, err := server.CheckCredentials(target.Login, target.Password)
	if err != nil {
		return ResultMsg{
			Error: errInternalError,
		}
	}

	c.Console.Server = server
	c.Console.Privilege = priv
	c.Console.Login = target.Login
	c.Console.InitMem()

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", server.Address)
	fmt.Fprintf(&b, "%s\n", server.Description)

	return ResultMsg{
		Output: b.String(),
	}
}
