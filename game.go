package main

import (
	"errors"
	"fmt"

	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
)

var (
	errInternalError      = errors.New("erreur interne")
	errServerNotFound     = errors.New("serveur introuvable")
	errServiceNotFound    = errors.New("serveur introuvable")
	errInvalidCommand     = errors.New("commande invalide")
	errMissingCommand     = errors.New("commande manquante")
	errMissingArgument    = errors.New("argument manquant")
	errInvalidCredentials = errors.New("identifiant ou mot de passe invalide")
	errInvalidLink        = errors.New("aucun service link ne porte ce nom")
	errInvalidDatabase    = errors.New("aucun service database ne porte ce nom")
	errNotConnected       = errors.New("la console n'est pas connectée")
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

func FindService[T any](g Game, address string, name string) (T, error) {
	var service T
	if err := g.Select(
		q.Eq("ServerAddress", address),
		q.Eq("Name", name),
	).First(&service); err != nil {
		if err == storm.ErrNotFound {
			return service, fmt.Errorf("%s : %w", name, errServiceNotFound)
		}

		// erreur interne
		fmt.Println(err)
		return service, errInternalError
	}

	return service, nil
}

func FindServices[T any](g Game, address string) ([]T, error) {
	var services []T
	if err := g.Find("ServerAddress", address, &services); err != nil {
		// erreur interne
		fmt.Println(err)
		return services, errInternalError
	}

	return services, nil
}

// Service regroupe les infos de base exposées par tous les services
type Service struct {
	ID            int    // ID du service (interne)
	ServerAddress string `storm:"index"` // ID du serveur sur lequel le service tourne
	Name          string `storm:"index"` // nom du service
	Description   string // description courte du service
	Privilege     int    // niveau de privilège requis pour utiliser le service
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
