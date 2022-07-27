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
	if !c.Console.IsConnected() {
		return ResultMsg{
			Cmd:   "registry search " + strings.Join(args, " "),
			Error: errNotConnected}
	}

	// par défaut afficher la liste de tous les registres
	regs := c.Console.Server.Registers

	if len(args) > 0 {
		// filter la liste des registres sur le préfixe
		regs = c.Console.Server.RegisterSearch(args[0])
	}

	b := strings.Builder{}
	tw := tw(&b)
	fmt.Fprintf(tw, "NAME\tSTATE\tDESCRIPTION\t\n")
	for _, r := range regs {
		if r.Restricted <= c.Privilege {
			fmt.Fprintf(tw, "%s\t%t\t%s\t\n", r.Name, r.State, r.Description)
		} else {
			fmt.Fprintf(tw, "%s\t%t\t%s\t\n", r.Name, r.State, "Accès restreint")
		}
	}
	tw.Flush()
	return ResultMsg{
		Cmd:    "registry search " + strings.Join(args, " "),
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
	if !c.Console.IsConnected() {
		return ResultMsg{
			Cmd:   "registry edit " + strings.Join(args, " "),
			Error: errNotConnected}
	}

	if len(args) < 1 {
		return ResultMsg{
			Cmd:   "registry edit " + strings.Join(args, " "),
			Error: fmt.Errorf("NAME : %w", errMissingArgument),
		}
	}

	name := args[0]
	reg, err := c.FindRegister(name)
	if err != nil {
		return ResultMsg{
			Cmd:   "registry edit " + strings.Join(args, " "),
			Error: fmt.Errorf("%s : %w", name, err),
		}
	}

	if reg.Restricted > c.Privilege {
		return ResultMsg{
			Cmd:   "registry edit " + strings.Join(args, " "),
			Error: errLowPrivilege,
		}
	}
	reg.State = !reg.State
	// Persistent: save new game state
	c.Game.Serialize()
	
	return ResultMsg{
		Cmd:    "registry edit " + strings.Join(args, " "),
		Output: fmt.Sprintf("registre %s est désormais sur l'état '%t'\n", reg.Name, reg.State),
	}
}
