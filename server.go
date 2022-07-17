package main

import (
	"strings"
	//"github.com/lithammer/fuzzysearch/fuzzy"
)

// Server représente un serveur sur le Net
type Server struct {
	// Addresse du serveur sur le réseau
	Address string

	// liste de codes d'accès valides pour se connecter au serveur
	Credentials []Cred

	// informations affichées lors de la connexion
	Description string

	// niveau de détection. plus il est élevé, plus vite on se fait repérer
	Detection float64

	// liste des liens fournis par le serveur
	Targets []Target

	// liste des données fournies par le serveur
	Entries []Entry

	// liste des registres fournis par le serveur
	Registers []Register
}

// Cred représente les droits d'accès d'un utilisateur à un serveur
type Cred struct {
	Login     string
	Password  string
	Privilege int
}

// CheckCredentials vérifie la validité de la paire login/password
// utilisé par la commande CONNECT
func (s *Server) CheckCredentials(login, password string) (int, error) {
	for _, c := range s.Credentials {
		if c.Login == login && c.Password == password {
			return c.Privilege, nil
		}
	}

	return 0, errInvalidCredentials
}

type Target struct {
	// adresse du serveur de destination
	Address string

	// description du lien
	Description string

	// niveau de privilège nécessaire pour utiliser ce target
	Restricted int

	// niveau de privilège obtenu après la connexion
	Privilege int
}

func (s *Server) FindTarget(address string) (Target, error) {
	for _, t := range s.Targets {
		if t.Address == address {
			return t, nil
		}
	}
	return Target{}, errInvalidArgument
}

// Entry est une entrée dans une base de données
type Entry struct {
	// clef unique
	ID string

	// mots-clefs utilisés pour la recherche
	Keywords []string

	// niveau de privilège requis
	Restricted int

	// titre de l'entrée
	Title string

	// contenu de l'entrée
	Content string
}

func (s *Server) DataSearch(keyword string) []Entry {
	result := make([]Entry, 0, len(s.Entries))
	for _, e := range s.Entries {
		if e.Match(keyword) {
			result = append(result, e)
		}
	}
	return result
}

func (s *Server) RegisterSearch(prefix string) []Register {
	result := make([]Register, 0, len(s.Registers))
	for _, r := range s.Registers {
		if r.Match(prefix) {
			result = append(result, r)
		}
	}
	return result
}

func (s *Server) FindEntry(id string) (Entry, error) {
	for _, e := range s.Entries {
		if e.ID == id {
			return e, nil
		}
	}
	return Entry{}, errInvalidArgument
}

func (s *Server) FindRegister(name string) (*Register, error) {
	for i, r := range s.Registers {
		if r.Name == name {
			return &s.Registers[i], nil
		}
	}
	return nil, errInvalidArgument
}

// Match détermine si l'entrée contient le mot-clef
func (e Entry) Match(keyword string) bool {
	//FIXME
	//keyword = normalize(keyword)

	for _, k := range e.Keywords {
		if k == keyword {
			return true
		}
	}

	return false
}

// Register représente registre mémoire qui peut être modifié pour contrôler quelque chose
type Register struct {
	Name        string
	State       bool
	Description string
}

func (r *Register) Match(prefix string) bool {
	return strings.HasPrefix(r.Name, prefix)
}
