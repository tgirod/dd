package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Link struct{}

func (l Link) ParseName() string {
	return "link"
}

func (l Link) ShortHelp() string {
	return "link -- affiche les liens disponibles ou suit un lien"
}

func (l Link) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(l.ShortHelp() + "\n")
	b.WriteString("USAGE\n")
	b.WriteString("  link\n")
	b.WriteString("    liste les liens disponibles\n")
	b.WriteString("  link <ID>\n")
	b.WriteString("    suit le lien ID\n")
	return b.String()
}

func (l Link) Run(c *Client, args []string) tea.Msg {
	if !c.Console.IsConnected() {
		return ResultMsg{
			Error: errNotConnected,
		}
	}

	if len(args) == 0 {
		// FIXME afficher la liste des liens
		// mention "accès restreint" quand l'utilisateur n'a pas les privilèges
		return ResultMsg{
			Output: "FIXME",
		}
	}

	// FIXME
	//id := args[0]
	// récupérer le lien correspondant
	// vérifier le niveau d'accréditation
	// récupérer le serveur correspondant
	// effectuer la connexion avec le serveur
	return ResultMsg{
		Output: "FIXME",
	}
}
