package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var Registry = Node{
	Name: "registry",
	Help: "manipuler les périphériques connectés au serveur",
	Sub: []Command{
		RegistrySearch{},
		RegistryEdit{},
	},
}

type RegistrySearch struct{}

func (r RegistrySearch) ParseName() string {
	return "search"
}

func (r RegistrySearch) ShortHelp() string {
	return "recherche des registres par nom et affiche leur état"
}

func (r RegistrySearch) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(r.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  registry search [NAME]\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  aucun -- liste tous les registres\n")
	b.WriteString("  NAME  -- nom complet ou partiel")
	return b.String()
}

func (r RegistrySearch) Run(c *Client, args []string) tea.Msg {
	cmd := fmt.Sprintf("registry search %s", strings.Join(args, " "))

	var name = ""
	if len(args) > 0 {
		name = args[0]
	}

	search, err := c.Console.RegistrySearch(name)
	if err != nil {
		return ResultMsg{
			Cmd:   cmd,
			Error: err,
		}
	}

	b := strings.Builder{}
	tw := tw(&b)
	fmt.Fprintf(tw, "NAME\tSTATE\tDESCRIPTION\t\n")
	for _, r := range search {
		fmt.Fprintf(tw, "%s\t%t\t%s\t\n", r.Name, r.State, r.Description)
	}
	tw.Flush()
	return ResultMsg{
		Cmd:    cmd,
		Output: b.String(),
	}
}

type RegistryEdit struct{}

func (r RegistryEdit) ParseName() string {
	return "edit"
}

func (r RegistryEdit) ShortHelp() string {
	return "change l'état d'un registre"
}

func (r RegistryEdit) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(r.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  registry edit <NAME>\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  NAME -- nom du registre")
	return b.String()
}

func (r RegistryEdit) Run(c *Client, args []string) tea.Msg {
	cmd := fmt.Sprintf("registry edit %s", strings.Join(args, " "))

	if len(args) < 1 {
		return ResultMsg{
			Cmd:   cmd,
			Error: fmt.Errorf("NAME : %w", errMissingArgument),
		}
	}

	name := args[0]
	if err := c.Console.RegistryEdit(name); err != nil {
		return ResultMsg{
			Cmd:   cmd,
			Error: err,
		}
	}

	return ResultMsg{
		Cmd:    cmd,
		Output: fmt.Sprintf("l'état du registre %s est changé\n", name),
	}
}
