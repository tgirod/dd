package main

import (
	"encoding/gob"
	"errors"

	"github.com/asdine/storm/v3"
	gc "github.com/asdine/storm/v3/codec/gob"
)

var (
	errInternalError   = errors.New("erreur interne")
	errServerNotFound  = errors.New("serveur introuvable")
	errInvalidCommand  = errors.New("commande invalide")
	errMissingCommand  = errors.New("commande manquante")
	errMissingArgument = errors.New("argument manquant")
	errInvalidLogin    = errors.New("identifiant invalide")
	errInvalidPassword = errors.New("mot de passe invalide")
	errInvalidLink     = errors.New("aucun service link ne porte ce nom")
	errInvalidDatabase = errors.New("aucun service database ne porte ce nom")
)

// Game contient l'état du jeu et les méthodes utiles pour en simplifier l'accès
type Game struct {
	*storm.DB
}

func NewGame(path string) (Game, error) {
	gob.Register(Node{})
	gob.Register(Connect{})
	db, err := storm.Open(path, storm.Codec(gc.Codec))
	return Game{db}, err
}

// Console représente le terminal depuis lequel le joueur accède au net
type Console struct {
	ID       int `storm:"id,increment"`
	Command      // commandes disponibles
	serverID int // identifiant du serveur auquel la console est connectée
}

func NewConsole() Console {
	return Console{
		Command: Node{
			Sub: []Command{
				Connect{},
			},
		},
	}
}

// Server représente un serveur sur le Net
type Server struct {
	// informations générales
	ID      int    // ID du serveur (interne)
	Address string `storm:"unique"` // Addresse du serveur sur le réseau

	// services
	Links     []Link
	Databases []Database

	// liste de codes d'accès valides pour se connecter au serveur
	Credentials map[string]struct {
		password  string
		privilege int
	}

	// TODO backdoors
}

// Connect vérifie que la tentative de connexion est valide
func (s Server) Connect(login, password string) (int, error) {
	cred, ok := s.Credentials[login]
	if !ok {
		return 0, errInvalidLogin
	}

	if cred.password != password {
		return 0, errInvalidPassword
	}

	return cred.privilege, nil
}

func (s Server) Link(name string) (int, error) {
	return 0, errInvalidLink
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

// Database est un service de base de données
type Database struct {
	Service `storm:"inline"`
	Entries []struct {
		keywords []string
		content  string
	}
}
