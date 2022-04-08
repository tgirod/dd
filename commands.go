package main

import (
	"fmt"
	"strings"

	"github.com/asdine/storm/v3"
	tea "github.com/charmbracelet/bubbletea"
)

// Command est l'interface implémentée par les commandes pour parser les
// arguments et générer l'aide
type Command interface {
	ParseName() string // nom de la commande pour le parsing
	ShortHelp() string // nom de la commande + ligne de description
	LongHelp() string  // aide complète
	Run(g Game, args []string) tea.Msg
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
	return fmt.Sprintf("%s -- %s", n.Name, n.Help)
}

func (n Node) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(n.ShortHelp() + "\n")
	b.WriteString("USAGE\n")
	fmt.Fprintf(&b, "\t%s <SOUS-COMMANDE>\n", n.Name)
	b.WriteString("SOUS-COMMANDES\n")
	for _, s := range n.Sub {
		b.WriteString("\t" + s.ShortHelp() + "\n")
	}
	return b.String()
}

// Parse exécute le parsing des arguments pour le noeud courant
func (n Node) Run(g Game, args []string) tea.Msg {
	if len(args) == 0 {
		return LogMsg{
			errMissingCommand,
			n.LongHelp(),
		}
	}

	match := n.Match(args[0])
	if len(match) == 0 {
		return LogMsg{
			errInvalidCommand,
			n.LongHelp(),
		}
	}

	// on retient la première commande qui a le bon préfixe
	return match[0].Run(g, args[1:])
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

// Connect établit la connexion à un serveur
type Connect struct{}

func (c Connect) ParseName() string {
	return "connect"
}

func (c Connect) ShortHelp() string {
	return "connect -- établit la connexion avec un serveur"
}

func (c Connect) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(c.ShortHelp() + "\n")
	b.WriteString("USAGE\n")
	b.WriteString("\tconnect <ADDRESS> <LOGIN> <PASSWORD>\n")
	b.WriteString("ARGUMENTS\n")
	b.WriteString("\t ADDRESS -- l'adresse du serveur sur le Net\n")
	b.WriteString("\t LOGIN -- identifiant de connexion\n")
	b.WriteString("\t PASSWORD -- mot de passe de connexion\n")
	return b.String()
}

func (c Connect) Run(g Game, args []string) tea.Msg {
	if len(args) < 3 {
		return LogMsg{
			errMissingArgument,
			c.LongHelp(),
		}
	}

	// récupérer les arguments
	address := args[0]
	login := args[1]
	password := args[2]

	// récupérer le serveur
	var server Server
	if err := g.One("Address", address, &server); err != nil {
		if err == storm.ErrNotFound {
			return LogMsg{
				err: fmt.Errorf("%s : %w", address, errServerNotFound),
			}
		} else {
			return LogMsg{
				err: err,
			}
		}
	}

	privilege, err := server.Connect(login, password)
	if err != nil {
		return LogMsg{
			err: fmt.Errorf("connexion impossible : %w", err),
		}
	}

	return ConnectMsg{server, privilege}
}

// ConnectMsg est retourné quand la connexion est une réussite
type ConnectMsg struct {
	Server        // infos sur le serveur
	Privilege int // niveau de privilège acquis
}
