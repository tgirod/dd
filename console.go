package main

import "fmt"

// Console représente le terminal depuis lequel le joueur accède au net
type Console struct {
	// identifiant unique pour la BDD
	ID int `storm:"id,increment"`

	// racine de l'arbre des commandes
	Node

	// identifiant dans le serveur actuel
	Login string

	// niveau de privilège dans le serveur actuel
	Privilege int

	// niveau d'alerte du serveur
	Alarm int

	// serveur auquel la console est connectée
	Server
}

func NewConsole(g Game) (Console, error) {
	// créer la console avec les commandes de base
	var console = Console{
		Node: Node{
			Sub: []Command{
				Connect{},
				Node{
					Name: "data",
					Help: "utiliser le service DATABASE pour rechercher des données",
					Sub: []Command{
						DataSearch{},
					},
				},
				Help{},
				Index{},
				Node{
					Name: "link",
					Help: "utiliser le service GATE pour accéder à un autre serveur",
					Sub: []Command{
						LinkList{},
						LinkConnect{},
					},
				},
				Quit{},
				Jack{},
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

func (c Console) HasAccess(restricted int) bool {
	return c.Privilege >= restricted
}