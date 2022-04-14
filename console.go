package main

import "fmt"

// Console représente le terminal depuis lequel le joueur accède au net
type Console struct {
	ID        int `storm:"id,increment"`
	Node          // commandes disponibles
	Privilege int // niveau de privilège
	Server        // serveur auquel la console est actuellement connectée
}

func NewConsole(g Game) (Console, error) {
	// créer la console avec les commandes de base
	var console = Console{
		Node: Node{
			Sub: []Command{
				Connect{},
				Help{},
				Index{},
				Link{},
				Quit{},
			},
		},
	}

	// sauver la console dans la BDD
	if err := g.Save(&console); err != nil {
		if err != nil {
			fmt.Println(err)
			return console, errInternalError
		}
	}

	return console, nil
}

func (c Console) IsConnected() bool {
	return c.Server.Address != ""
}