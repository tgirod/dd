package main

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Jack struct{}

func (j Jack) ParseName() string {
	return "jack"
}

func (j Jack) ShortHelp() string {
	return "jack\tforce l'accès à un lien"
}

func (j Jack) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(j.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  jack [ID]\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  aucun -- liste les liens disponibles\n")
	b.WriteString("  ID    -- force l'accès au lien ID\n\n")
	return b.String()
}

func (j Jack) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{
			Cmd:   "jack " + strings.Join(args, " "),
			Error: errNotConnected,
		}
	}

	if len(args) == 0 {
		// lister les liens disponibles
		b := strings.Builder{}
		tw := tw(&b)
		fmt.Fprintf(tw, "ID\tDESCRIPTION\t\n")
		for i, t := range c.Server.Targets {
			fmt.Fprintf(tw, "%d\t%s\t\n", i, t.Description)
		}
		tw.Flush()

		return ResultMsg{
			Cmd:    "jack",
			Output: b.String(),
		}
	}

	// récupérer le lien
	id, err := strconv.Atoi(args[0])
	if err != nil || id < 0 || id >= len(c.Server.Targets) {
		return ResultMsg{
			Cmd:   "jack " + strings.Join(args, " "),
			Error: fmt.Errorf("ID : %w", errInvalidArgument),
		}
	}
	target := c.Server.Targets[id]

	// récupérer le serveur correspondant
	server, err := c.Game.FindServer(target.Address)
	if err != nil {
		return ResultMsg{
			Cmd:   "jack " + strings.Join(args, " "),
			Error: fmt.Errorf("%s : %w", target.Address, err),
		}
	}

	co := c.Console
	co.Server = server
	co.Login = "illegal"
	co.Privilege = 1
	co.Alert++
	co.InitMem()

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", server.Address)
	fmt.Fprintf(&b, "%s\n", server.Description)

	return ResultMsg{
		Cmd:    "jack " + strings.Join(args, " "),
		Output: b.String(),
	}
}
