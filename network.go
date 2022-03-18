package main

import "errors"

var (
	errNotImplemented  = errors.New("pas encore implémenté")
	errNotConnected    = errors.New("aucune connexion active")
	errDNIRequired     = errors.New("Interface Neurale Directe requise")
	errServerNotFound  = errors.New("serveur introuvable")
	errAccessDenied    = errors.New("mot de passe invalide")
	errServiceNotFound = errors.New("service introuvable")
	errCannotConnect   = errors.New("impossible de se connecter")
	errCannotSearch    = errors.New("impossible de lancer une recherche")
	errMissingParam    = errors.New("paramètre manquant")
	errInvalidParam    = errors.New("paramètre invalide")
	errCommandNotFound = errors.New("commande introuvable")
	errForceTooLow     = errors.New("force de l'attaque insuffisante")
)

type Network struct {
	Servers []*Server
}

func (n *Network) FindServer(address string) (*Server, error) {
	for _, s := range n.Servers {
		if s.Address == address {
			return s, nil
		}
	}
	return nil, errServerNotFound
}
