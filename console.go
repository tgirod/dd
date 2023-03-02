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

	// identifiant dans le serveur actuel
	Login string

	// niveau de privilège dans le serveur actuel
	Privilege int

	// l'alerte est-elle activée ?
	Alert bool

	// durée avant la déconnexion forcée
	Countdown time.Duration

	// zones mémoires disponibles pour une évasion
	Mem map[string]bool

	// interface neurale directe
	DNI bool

	// serveur auquel la console est connectée
	*Server

	// Pile de serveurs visités lors de cette connexion
	History Stack

	// Forum
	Forum
}

var Hack = map[string]Command{
	"yyqz": Jack{},
	"hfed": Rise{},
	"zfcq": Evade{},
}

func NewConsole() *Console {
	return &Console{
		Node: Node{
			Sub: []Command{
				Back{},
				Connect{},
				Data,
				ForumCmd,
				Help{},
				Identify{},
				Index{},
				Link{},
				Load{},
				Plug{},
				Quit{},
				Registry,
			},
		},
	}
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
