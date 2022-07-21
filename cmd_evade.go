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
	return "evade\tmanoeuvre d'évasion pour gagner un peu de temps"
}

func (e Evade) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(e.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  evade [ZONE]\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  aucun -- lister les zones mémoire disponibles\n")
	b.WriteString("  ZONE  -- évasion vers la zone mémoire\n")
	return b.String()
}

func (e Evade) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{
			Cmd:   "evade" + strings.Join(args, " "),
			Error: errNotConnected,
		}
	}

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
			Cmd:    "evade",
			Output: b.String(),
		}
	}

	addr := args[0]
	available, exist := c.Console.Mem[addr]
	if !exist {
		return ResultMsg{
			Cmd:   "evade" + strings.Join(args, " "),
			Error: fmt.Errorf("%s : %w", addr, errMemNotFound),
		}
	}

	if !available {
		return ResultMsg{
			Cmd:   "evade" + strings.Join(args, " "),
			Error: fmt.Errorf("%s : %w", addr, errMemUnavailable),
		}
	}

	// évasion effectuée
	c.Console.Mem[addr] = false
	c.Console.Alert = 1

	return ResultMsg{
		Cmd:    "evade " + strings.Join(args, " "),
		Output: "Evasion effectuée",
	}
}
