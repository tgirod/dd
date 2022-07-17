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
	return "evade -- manoeuvre d'évasion pour gagner un peu de temps"
}

func (e Evade) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(e.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  evade [MEM]\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  aucun -- liste les zones mémoire disponibles\n")
	b.WriteString("  MEM   -- évasion vers la zone mémoire\n")
	return b.String()
}

func (e Evade) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{Error: errNotConnected}
	}

	// afficher la liste des zones mémoires disponibles
	if len(args) == 0 {
		b := strings.Builder{}
		b.WriteString("ZONES MEMOIRES POUR EVASION\n")
		for addr, available := range c.Console.Mem {
			if !available {
				fmt.Fprintf(&b, "  %s  INDISPONIBLE\n", addr)
			} else {
				fmt.Fprintf(&b, "  %s  OK\n", addr)
			}
		}
		return ResultMsg{
			Output: b.String(),
		}
	}

	addr := args[0]
	available, exist := c.Console.Mem[addr]
	if !exist {
		return ResultMsg{
			Error: fmt.Errorf("%s : %w", addr, errMemNotFound),
		}
	}

	if !available {
		return ResultMsg{
			Error: fmt.Errorf("%s : %w", addr, errMemUnavailable),
		}
	}

	// évasion effectuée
	c.Console.Mem[addr] = false
	c.Console.Alert = 1

	return ResultMsg{
		Output: "Evasion effectuée",
	}
}