package main

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Link struct{}

func (l Link) ParseName() string {
	return "link"
}

func (l Link) ShortHelp() string {
	return "affiche les liens disponibles ou suit un lien"
}

func (l Link) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(l.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  link [ID]\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  aucun -- liste les liens disponibles\n")
	b.WriteString("  ID    -- suit le lien ID")
	return b.String()
}

func (l Link) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{
			Cmd:   "link " + strings.Join(args, " "),
			Error: errNotConnected,
		}
	}

	if len(args) == 0 {
		// lister les liens disponibles
		b := strings.Builder{}
		tw := tw(&b)
		fmt.Fprintf(tw, "ID\tDESCRIPTION\t\n")
		for i, t := range c.Server.Targets {
			if c.Console.Privilege >= t.Restricted {
				fmt.Fprintf(tw, "%d\t%s\t\n", i, t.Description)
			} else {
				fmt.Fprintf(tw, "\t%d\t%s\t\n", i, "Accès restreint")
			}
		}
		tw.Flush()

		return ResultMsg{
			Cmd:    "link " + strings.Join(args, " "),
			Output: b.String(),
		}
	}

	// récupérer le lien
	id, err := strconv.Atoi(args[0])
	if err != nil || id < 0 || id >= len(c.Server.Targets) {
		return ResultMsg{
			Cmd:   "link " + strings.Join(args, " "),
			Error: fmt.Errorf("ID : %w", errInvalidArgument),
		}
	}
	target := c.Server.Targets[id]

	// vérifier le niveau de privilège
	if c.Console.Privilege < target.Restricted {
		return ResultMsg{
			Cmd:   "link " + strings.Join(args, " "),
			Error: errLowPrivilege,
		}
	}

	// récupérer le serveur correspondant
	server, err := c.Game.FindServer(target.Address)
	if err != nil {
		return ResultMsg{
			Cmd:   "link " + strings.Join(args, " "),
			Error: fmt.Errorf("%s : %w", target.Address, err),
		}
	}

	// effectuer la connexion avec le serveur
	priv, err := server.CheckCredentials(target.Login, target.Password)
	if err != nil {
		return ResultMsg{
			Cmd:   "link " + strings.Join(args, " "),
			Error: errInternalError,
		}
	}

	c.Console.Server = server
	c.Console.Privilege = priv
	c.Console.Login = target.Login
	c.Console.InitMem()
	c.Console.History.Push(target)

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", server.Address)
	fmt.Fprintf(&b, "%s\n", server.Description)

	return ResultMsg{
		Cmd:    "link " + strings.Join(args, " "),
		Output: b.String(),
	}
}
