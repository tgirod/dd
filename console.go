package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/mattn/go-shellwords"
)

const MAX_RESULTS int = 10

// Console représente le terminal depuis lequel le joueur accède au net
type Console struct {
	ID int

	// racine de l'arbre des commandes
	Cmd

	// identité active sur la console
	*Identity

	// compte utilisateur sur le serveur courant
	*Account

	// l'alerte est-elle activée ?
	Alert bool

	// durée avant la déconnexion forcée
	Countdown time.Duration

	// zones mémoires disponibles pour une évasion
	Mem map[string]bool

	// interface neurale directe
	DNI bool

	// Pile de serveurs visités lors de cette connexion
	History Stack

	// liste des dernières commandes évaluées
	Results []Result

	// serveur auquel la console est connectée
	*Server

	// état interne du jeu
	*Network
}

type Result struct {
	Prompt string
	Error  error
	Output string
}

var Hack = map[string]Cmd{
	"jack":  jack,
	"evade": evade,
	"door":  door,
}

var baseCmds = Cmd{
	SubCmds: []Cmd{
		back,
		yes,
		connect,
		data,
		help,
		identify,
		index,
		link,
		load,
		plug,
		quit,
		registry,
		message,
		forum,
	},
}

func NewConsole(net *Network) *Console {
	return &Console{
		Cmd:     baseCmds,
		Network: net,
	}
}

func (c *Console) Parse(prompt string) any {
	args, err := shellwords.Parse(prompt)
	if err != nil {
		return Result{
			Prompt: prompt,
			Error:  err,
			Output: "",
		}
	}
	// args := strings.Fields(prompt)
	ctx := Context{
		Console: c,
		Args:    args,
		Cmd:     c.Cmd,
	}
	return ctx.Parse()
}

func (c *Console) connect(address string) error {
	var err error
	var server *Server
	var account *Account

	// récupérer le serveur
	if server, err = c.FindServer(address); err != nil {
		return fmt.Errorf("%s : %w", address, err)
	}

	// vérifier que l'utilisateur a le droit de se connecter
	login := ""
	if c.Identity != nil {
		login = c.Identity.Login
	}
	if account, err = server.CheckAccount(login); err != nil {
		return err
	}
	c.Account = account

	// enregistrer le nouveau serveur
	c.Server = server

	if c.Account != nil && c.Account.Backdoor {
		c.Server.RemoveAccount(c.Account.Login)
		c.Network.RemoveIdentity(c.Account.Login)
	}
	c.InitMem()
	return nil
}

func (c *Console) InitMem() {
	c.Mem = make(map[string]bool)
	for i := 0; i < 5; i++ {
		addr := fmt.Sprintf("%08x", rand.Uint32())
		c.Mem[addr] = true
	}
}

func (c *Console) Disconnect() {
	c.Server = nil
	c.Admin = false
	c.Alert = false
	c.History.Clear()
	c.Cmd = baseCmds

	// affichage par défaut
	eval := Result{
		Output: "coupure de la connexion au réseau.",
	}

	if c.DNI {
		eval.Output = `
			     DUMPSHOCK !!!!
                     _____
                    /     \
                   | () () |
                    \  ^  /
                     |||||
                     |||||

			PERDS UN POINT DE VIE

coupure de la connexion au réseau.`
	}

	c.AddResult(eval)
}

func (c *Console) StartAlert() {
	if !c.Alert {
		c.Alert = true
		c.Countdown = c.Server.Scan
	} else if c.Server.Scan < c.Countdown {
		c.Countdown = c.Server.Scan
	}
}

func (c *Console) Identify(login, password string) error {
	identity, err := c.CheckIdentity(login, password)
	if err != nil {
		return err
	}
	c.Identity = identity

	// si on est connecté à un serveur, on tente d'accéder au compte utilisateur
	if c.Server != nil {
		if account, err := c.CheckAccount(login); err == nil {
			c.Account = account
		}
	}

	return nil
}

func (c *Console) AddResult(o Result) {
	c.Results = append(c.Results, o)
	if len(c.Results) > MAX_RESULTS {
		c.Results = c.Results[len(c.Results)-MAX_RESULTS : len(c.Results)]
	}
}

// MessageNew affiche les messages non lus
func (c *Console) MessageNew() {
	b := strings.Builder{}
	tw := tw(&b)

	fmt.Fprintf(tw, "liste des messages non lus :\n")
	for i, m := range c.Messages {
		if !m.Opened {
			fmt.Fprintf(tw, "%d\t%s\t\n", i, m.Subject)
		}
	}
	tw.Flush()

	c.AddResult(Result{
		Prompt: "message new",
		Output: b.String(),
	})
}

// func (c *Console) MessageReply(index int) (MessageSendMsg, error) {
// 	if c.Identity == nil || index < 0 || index >= len(c.Identity.Messages) {
// 		return MessageSendMsg{}, errMessageNotFound
// 	}

// 	msg := c.Identity.Messages[index]
// 	return MessageSendMsg{
// 		Recipient: msg.Sender,
// 		Subject:   fmt.Sprintf("Re: %s", msg.Subject),
// 	}, nil
// }

func (c *Console) Delay() time.Duration {
	if c.DNI {
		return time.Second * DNISpeed
	} else {
		return time.Second
	}
}

func (c *Console) Connect(address string, force bool) error {
	server, err := c.FindServer(address)
	if err != nil {
		return err
	}

	login := ""
	if c.Identity != nil {
		login = c.Identity.Login
	}

	if !force {
		account, err := server.CheckAccount(login)
		if err != nil {
			return err
		}

		c.Account = account
	}

	c.Server = server

	if c.Account != nil && c.Account.Backdoor {
		c.RemoveAccount(login)
		c.RemoveIdentity(login)
	}

	c.InitMem()

	return nil
}
