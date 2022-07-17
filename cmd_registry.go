package main

import (
	"fmt"
	"strings"
	"text/tabwriter"

	tea "github.com/charmbracelet/bubbletea"
)

var Registry = Node{
	Name: "registry",
	Help: "manipuler les périphériques connectés au serveur",
	Sub: []Command{
		RegistryView{},
		RegistryEdit{},
	},
}

type RegistryView struct{}

func (r RegistryView) ParseName() string {
	return "view"
}

func (r RegistryView) ShortHelp() string {
	return "view -- affiche les registres correspondant à la recherche"
}

func (r RegistryView) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(r.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  registry view [PREFIX]\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  aucun   -- liste tous les registres\n")
	b.WriteString("  PREFIX  -- premières lettres du nom du registre\n")
	return b.String()
}

func (r RegistryView) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{Error: errNotConnected}
	}

	// par défaut afficher la liste de tous les registres
	regs := c.Console.Server.Registers

	if len(args) > 0 {
		// filter la liste des registres sur le préfixe
		regs = c.Console.Server.RegisterSearch(args[0])
	}

	b := strings.Builder{}
	tw := tabwriter.NewWriter(&b, 8, 1, 2, ' ', 0)
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
		Output: b.String(),
	}
}

type RegistryEdit struct{}

func (r RegistryEdit) ParseName() string {
	return "edit"
}

func (r RegistryEdit) ShortHelp() string {
	return "edit -- change l'état d'un registre"
}

func (r RegistryEdit) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(r.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  registry edit <NAME>\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  NAME -- nom du registre\n")
	return b.String()
}

func (r RegistryEdit) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{Error: errNotConnected}
	}

	if len(args) < 1 {
		return ResultMsg{
			Error: fmt.Errorf("NAME : %w", errMissingArgument),
		}
	}

	name := args[0]
	reg, err := c.FindRegister(name)
	if err != nil {
		return ResultMsg{
			Error: fmt.Errorf("%s : %w", name, err),
		}
	}

	if reg.Restricted > c.Privilege {
		return ResultMsg{
			Error: errLowPrivilege,
		}
	}
	reg.State = !reg.State

	return ResultMsg{
		Output: fmt.Sprintf("registre %s est désormais sur l'état '%t'\n", reg.Name, reg.State),
	}
}
