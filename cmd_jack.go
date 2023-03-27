package main

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Jack struct{}

type JackMsg struct {
	Id int
}

func (j Jack) ParseName() string {
	return "jack"
}

func (j Jack) ShortHelp() string {
	return "force l'accès à un lien"
}

func (j Jack) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(j.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  jack <ID>\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  ID    -- force l'accès au lien ID")
	return b.String()
}

func (j Jack) Run(c *Client, args []string) tea.Msg {
	cmd := fmt.Sprintf("jack %s", strings.Join(args, " "))
	if len(args) == 0 {
		return ResultMsg{
			Cmd:   cmd,
			Error: fmt.Errorf("ID : %w", errMissingArgument),
		}
	}

	// récupérer le lien
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return ResultMsg{
			Cmd:   cmd,
			Error: fmt.Errorf("ID : %w", errInvalidArgument),
		}
	}

	return JackMsg{id}
}
