package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Rise struct{}

func (r Rise) ParseName() string {
	return "rise"
}

func (r Rise) ShortHelp() string {
	return "rise -- augmente les privilèges d'un niveau"
}

func (r Rise) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(r.ShortHelp() + "\n")
	b.WriteString("USAGE\n")
	b.WriteString("  rise\n")
	return b.String()
}

func (r Rise) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{Error: errNotConnected}
	}

	c.Console.Privilege++
	c.Console.Alert++
	c.Console.Illegal()

	return ResultMsg{
		Output: fmt.Sprintf("niveau de privilège augmenté à %d", c.Console.Privilege),
	}
}
