package main

import (
	"fmt"
	//"github.com/lithammer/fuzzysearch/fuzzy"
)

// Game contient l'état du jeu et les méthodes utiles pour en simplifier l'accès
type Game struct {
	Network []Server
}

func (g Game) FindServer(address string) (*Server, error) {
	for _, server := range g.Network {
		if server.Address == address {
			return &server, nil
		}
	}
	return nil, fmt.Errorf("%s : %w", address, errServerNotFound)
}
