package main

import (
	"errors"
	//"fmt"
	"github.com/asdine/storm/v3"
)

var (
	errServerNotFound = errors.New("serveur introuvable")
)

// Game contient l'état du jeu et les méthodes utiles pour en simplifier l'accès
type Game struct {
	*storm.DB
}

// Server regroupe les infos concernant un serveur sur le Net
type Server struct {
	ID      int    // ID du serveur (interne)
	Address string `storm:"unique"` // Addresse du serveur sur le réseau

	// liste de codes d'accès valides pour se connecter au serveur
	Credentials map[string]struct {
		password  string
		privilege int
	}

	// TODO backdoors
}

// Service regroupe les infos de base exposées par tous les services
type Service struct {
	ID        int    // ID du service (interne)
	ServerID  int    // ID du serveur sur lequel le service tourne
	Name      string // nom du service
	Privilege int    // niveau de privilège requis pour utiliser le service
}

// Link est un service permettant de se connecter ailleurs
type Link struct {
	Service   `storm:"inline"`
	TargetID  int // ID du serveur distant
	Privilege int // niveau de privilège une fois connecté
}

// Data est un service de base de données
type Data struct {
	Service `storm:"inline"`
	Entries []struct {
		keywords []string
		content  string
	}
}
