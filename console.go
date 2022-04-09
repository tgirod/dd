package main

import "fmt"

// Console représente le terminal depuis lequel le joueur accède au net
type Console struct {
	ID        int `storm:"id,increment"`
	Node          // commandes disponibles
	ServerID  int // identifiant du serveur auquel la console est connectée
	Privilege int // niveau de privilège
}

func NewConsole(g Game) (Console, error) {
	// créer la console avec les commandes de base
	var console = Console{
		Node: Node{
			Sub: []Command{
				Connect{},
				Help{},
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
