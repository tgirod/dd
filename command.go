package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Command est l'interface implémentée par les commandes pour parser les
// arguments et générer l'aide
type Command interface {
	ParseName() string // nom de la commande pour le parsing
	ShortHelp() string // nom de la commande + ligne de description
	LongHelp() string  // aide complète
	Run(c *Client, args []string) tea.Msg
}

type Context struct {
	Game    // référence à l'état du jeu
	Console // la console d'ou provient la commande
}

// Node est un noeud intermédiaire dans l'arbre de commandes
type Node struct {
	Name string
	Help string
	Sub  []Command
}

func (n Node) ParseName() string {
	return n.Name
}

func (n Node) ShortHelp() string {
	return fmt.Sprintf("%s\t%s", n.Name, n.Help)
}

func (n Node) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(n.ShortHelp() + "\n")
	b.WriteString("USAGE\n")
	fmt.Fprintf(&b, "  %s <SOUS-COMMANDE>\n", n.Name)
	b.WriteString("SOUS-COMMANDES\n")
	for _, s := range n.Sub {
		b.WriteString("  " + s.ShortHelp() + "\n")
	}
	return b.String()
}

// Parse exécute le parsing des arguments pour le noeud courant
func (n Node) Run(c *Client, args []string) tea.Msg {
	if len(args) == 0 {
		return ResultMsg{
			Error:  errMissingCommand,
			Output: n.LongHelp(),
		}
	}

	match := n.Match(args[0])
	if len(match) == 0 {
		return ResultMsg{
			Error:  fmt.Errorf("%s : %w", args[0], errInvalidCommand),
			Output: n.LongHelp(),
		}
	}

	// on retient la première commande qui a le bon préfixe
	return match[0].Run(c, args[1:])
}

// Match retourne la liste des sous-commandes correspondant au préfixe
func (n Node) Match(prefix string) []Command {
	sub := make([]Command, 0, len(n.Sub))
	for _, s := range n.Sub {
		if strings.HasPrefix(s.ParseName(), prefix) {
			sub = append(sub, s)
		}
	}
	return sub
}
