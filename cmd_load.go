package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Load struct{}

func (l Load) ParseName() string {
	return "load"
}

func (l Load) ShortHelp() string {
	return "load -- charge une nouvelle commande"
}

func (l Load) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(l.ShortHelp() + "\n")
	b.WriteString("USAGE\n")
	b.WriteString("  load <CODE>\n")
	b.WriteString("ARGUMENTS\n")
	b.WriteString("  CODE -- code de la commande\n")
	return b.String()
}

func (l Load) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{Error: errNotConnected}
	}

	if len(args) < 1 {
		return ResultMsg{
			Error:  fmt.Errorf("CODE : %w", errMissingArgument),
			Output: l.LongHelp(),
		}
	}

	code := args[0]
	cmd, ok := Hack[code]
	if !ok {
		return ResultMsg{
			Error: fmt.Errorf("%s : %w", code, errInvalidArgument),
		}
	}
	c.Console.Node.Sub = append(c.Console.Node.Sub, cmd)

	return ResultMsg{
		Output: fmt.Sprintf("%s : commande chargée", cmd.ParseName()),
	}
}
