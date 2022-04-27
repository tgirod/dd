package main

import (
	"errors"
	"fmt"

	"github.com/asdine/storm/v3"
)

var (
	errInternalError      = errors.New("erreur interne")
	errServerNotFound     = errors.New("serveur introuvable")
	errServiceNotFound    = errors.New("serveur introuvable")
	errInvalidCommand     = errors.New("commande invalide")
	errMissingCommand     = errors.New("commande manquante")
	errMissingArgument    = errors.New("argument manquant")
	errInvalidArgument    = errors.New("argument invalide")
	errInvalidCredentials = errors.New("identifiant ou mot de passe invalide")
	errGateNotFound       = errors.New("service gate introuvable")
	errDatabaseNotFound   = errors.New("service database introuvable")
	errNotConnected       = errors.New("la console n'est pas connectée")
	errLowPrivilege       = errors.New("niveau de privilège insuffisant")
)

// Game contient l'état du jeu et les méthodes utiles pour en simplifier l'accès
type Game struct {
	*storm.DB
}

// Server représente un serveur sur le Net
type Server struct {
	// informations générales
	Address string `storm:"id"` // Addresse du serveur sur le réseau

	// liste de codes d'accès valides pour se connecter au serveur
	Credentials []Cred

	// niveau de détection. plus il est élevé, plus vite on se fait repérer
	Detection float64

	// les services fournis par le serveur
	Gate
	Database
}

type Cred struct {
	Login     string
	Password  string
	Privilege int
}

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

func (g Game) FindServer(address string) (Server, error) {
	var server Server
	if err := g.One("Address", address, &server); err != nil {
		if err == storm.ErrNotFound {
			return server, fmt.Errorf("%s : %w", address, errServerNotFound)
		}

		// erreur interne
		fmt.Println(err)
		return server, errInternalError
	}

	return server, nil
}

func (g Game) CreateConsole() (Console, error) {
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
				Rise{},
			},
		},
	}

	// sauver la console dans la BDD
	if err := g.Save(&console); err != nil {
		if err != nil {
			fmt.Println(err)
			return console, err
		}
	}

	// retourner la console créée pour le client
	return console, nil
}

func (g Game) UpdateConsole(console Console) error {
	// sauver la console dans la BDD
	if err := g.Save(&console); err != nil {
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func (g Game) DeleteConsole(console Console) error {
	if err := g.DeleteStruct(&console); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s Server) CheckCredentials(login, password string) (int, error) {
	for _, c := range s.Credentials {
		if c.Login == login && c.Password == password {
			return c.Privilege, nil
		}
	}

	return 0, errInvalidCredentials
}

func (c Console) IsConnected() bool {
	return c.Server.Address != ""
}
