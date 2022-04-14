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
	errInvalidCredentials = errors.New("identifiant ou mot de passe invalide")
	errGateNotFound       = errors.New("service gate introuvable")
	errDataNotFound       = errors.New("service data introuvable")
	errNotConnected       = errors.New("la console n'est pas connectée")
	errLowPrivilege       = errors.New("niveau de privilège insuffisant")
)

// Game contient l'état du jeu et les méthodes utiles pour en simplifier l'accès
type Game struct {
	*storm.DB
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

// Service regroupe les infos de base exposées par tous les services
type Service struct {
	ID          int    `storm:"increment"` // ID du service (interne)
	Name        string `storm:"index"`     // nom du service
	Description string // description courte du service
	Restricted  int    // niveau de privilège requis pour utiliser le service
}

// Gate est un service permettant de se connecter ailleurs
type Gate struct {
	Service       `storm:"inline"`
	TargetAddress string // ID du serveur distant
	Privilege     int    // niveau de privilège une fois connecté
}

// Database est un service de base de données
type Database struct {
	Service `storm:"inline"`
	Entries []struct {
		keywords []string
		content  string
	}
}
