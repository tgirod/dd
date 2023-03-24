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
	return "charge une nouvelle commande"
}

func (l Load) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(l.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  load <CODE>\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  CODE -- code de la commande")
	return b.String()
}

func (l Load) Run(c *Client, args []string) tea.Msg {
	cmd := fmt.Sprintf("load %s", strings.Join(args, " "))
	if len(args) < 1 {
		return ResultMsg{
			Cmd:    cmd,
			Error:  fmt.Errorf("CODE : %w", errMissingArgument),
			Output: l.LongHelp(),
		}
	}

	code := args[0]
	if err := c.Load(code); err != nil {
		return ResultMsg{
			Cmd:    cmd,
			Error:  err,
			Output: l.LongHelp(),
		}
	}

	cmds := c.Console.Node.Sub
	name := cmds[len(cmds)-1].ParseName()
	return ResultMsg{
		Cmd:    "load " + strings.Join(args, " "),
		Output: fmt.Sprintf("%s : commande chargée", name),
	}
}
