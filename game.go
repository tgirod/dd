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
