package main

import (
	"fmt"
	"math/rand"
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

	// niveau d'alerte du serveur
	Alert int

	// zones mémoires disponibles pour une évasion
	Mem map[string]bool

	// interface neurale directe
	DNI bool

	// serveur auquel la console est connectée
	*Server
}

var Hack = map[string]Command{
	"jack":  Jack{},
	"rise":  Rise{},
	"evade": Evade{},
}

func NewConsole() *Console {
	return &Console{
		Node: Node{
			Sub: []Command{
				Connect{},
				Node{
					Name: "data",
					Help: "effectuer une recherche sur le serveur",
					Sub: []Command{
						DataSearch{},
						DataView{},
					},
				},
				Help{},
				Index{},
				Link{},
				Load{},
				Plug{},
				Quit{},
				Node{
					Name: "registry",
					Help: "manipuler les périphériques connectés au serveur",
					Sub: []Command{
						RegistryView{},
						RegistryEdit{},
					},
				},
			},
		},
	}
}

func (c *Console) IsConnected() bool {
	return c.Server != nil
}

// Illegal est une méthode appelée à chaque fois qu'une commande illégale est utilisée
func (c *Console) Illegal() {
	if len(c.Mem) == 0 {
		c.InitMem()
	}
}

func (c *Console) InitMem() {
	c.Mem = make(map[string]bool)
	for i := 0; i < 5; i++ {
		addr := fmt.Sprintf("%08x", rand.Uint32())
		c.Mem[addr] = true
	}
}
