package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Console représente le terminal depuis lequel le joueur accède au net
type Console struct {
	ID int

	// racine de l'arbre des commandes
	Node

	// identité active sur la console
	Identity string

	// identifiant dans le serveur actuel
	Login string

	// admin dans le serveur actuel ?
	Admin bool

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

	// serveur auquel la console est connectée
	*Server

	// état interne du jeu
	*Game
}

var Hack = map[string]Command{
	"yyqz": Jack{},
	"zfcq": Evade{},
}

var baseCmds = []Command{
	Back{},
	Connect{},
	Data,
	Help{},
	Identify{},
	Index{},
	Link{},
	Load{},
	Plug{},
	Quit{},
	Registry,
	Pop{},
}

func NewConsole(game *Game) *Console {
	return &Console{
		Node: Node{
			Sub: baseCmds,
		},
		Game: game,
	}
}

func (c *Console) Connect(s *Server, admin bool) {
	c.Admin = admin
	c.Server = s
	c.InitMem()
}

func (c *Console) IsConnected() bool {
	return c.Server != nil
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
	c.Login = ""
	c.Admin = false
	c.Alert = false
	c.History.Clear()
	c.Node.Sub = baseCmds
}
