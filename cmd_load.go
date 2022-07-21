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
	return "load\tcharge une nouvelle commande"
}

func (l Load) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(l.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  load <CODE>\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  CODE -- code de la commande\n\n")
	return b.String()
}

func (l Load) Run(c *Client, args []string) tea.Msg {
	if len(args) < 1 {
		return ResultMsg{
			Cmd:    "load " + strings.Join(args, " "),
			Error:  fmt.Errorf("CODE : %w", errMissingArgument),
			Output: l.LongHelp(),
		}
	}

	code := args[0]
	cmd, ok := Hack[code]
	if !ok {
		return ResultMsg{
			Cmd:   "load " + strings.Join(args, " "),
			Error: fmt.Errorf("%s : %w", code, errInvalidArgument),
		}
	}
	c.Console.Node.Sub = append(c.Console.Node.Sub, cmd)

	return ResultMsg{
		Cmd:    "load " + strings.Join(args, " "),
		Output: fmt.Sprintf("%s : commande chargée", cmd.ParseName()),
	}
}
