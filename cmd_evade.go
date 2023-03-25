package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Evade struct{}

func (e Evade) ParseName() string {
	return "evade"
}

func (e Evade) ShortHelp() string {
	return "manoeuvre d'évasion pour gagner un peu de temps"
}

func (e Evade) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(e.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  evade [ZONE]\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  aucun -- lister les zones mémoire disponibles\n")
	b.WriteString("  ZONE  -- évasion vers la zone mémoire")
	return b.String()
}

func (e Evade) Run(c *Client, args []string) tea.Msg {
	cmd := fmt.Sprintf("evade %s", strings.Join(args, " "))

	// afficher la liste des zones mémoires disponibles
	if len(args) == 0 {
		b := strings.Builder{}
		tw := tw(&b)
		fmt.Fprintf(tw, "ZONE\tDISPONIBILITE\t\n")
		for addr, available := range c.Console.Mem {
			if !available {
				fmt.Fprintf(tw, "%s\t%s\t\n", addr, "INDISPONIBLE")
			} else {
				fmt.Fprintf(tw, "%s\t%s\t\n", addr, "OK")
			}
		}
		tw.Flush()

		return ResultMsg{
			Cmd:     cmd,
			Output:  b.String(),
			Illegal: true,
		}
	}

	zone := args[0]
	if err := c.Console.Evade(zone); err != nil {
		return ResultMsg{
			Cmd:   cmd,
			Error: err,
		}
	}

	return ResultMsg{
		Cmd:     cmd,
		Output:  "Evasion effectuée",
		Illegal: true,
	}
}
