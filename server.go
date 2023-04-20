package main

import (
	"fmt"
	"time"
	"unicode"

	"github.com/lithammer/fuzzysearch/fuzzy"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

const (
	PUBLIC_PRIVILEGE int = 1
)

// Server représente un serveur sur le Net
type Server struct {
	// Addresse du serveur sur le réseau
	Address string

	// ce serveur accepte-t-il des connexions anonymes ?
	Public bool

	// liste des comptes utilisateurs enregistrés
	Accounts []Account

	// informations affichées lors de la connexion
	Description string

	// durée du scan avant de se faire repérer par le serveur
	Scan time.Duration

	// liste des liens fournis par le serveur
	Links []Link

	// liste des données fournies par le serveur
	Entries []Entry

	// liste des registres fournis par le serveur
	Registers []Register
}

// Account représente un compte utilisateur sur un serveur
type Account struct {
	Login    string
	Admin    bool
	Backdoor bool
}

func (s *Server) CheckAccount(login string) (*Account, error) {
	// cherche un compte utilisateur valide
	for i, a := range s.Accounts {
		if a.Login == login {
			return &s.Accounts[i], nil
		}
	}

	// si le serveur est public, autoriser l'accès quoi qu'il arrive
	if s.Public {
		return nil, nil
	}

	return nil, fmt.Errorf("%s : %w", login, errInvalidIdentity)
}

func (s *Server) RemoveAccount(login string) {
	for i, a := range s.Accounts {
		if a.Login == login {
			// retirer la backdoor après usage
			last := len(s.Accounts) - 1
			s.Accounts[i] = s.Accounts[last]
			s.Accounts = s.Accounts[:last]
			return
		}
	}
}

type Link struct {
	// adresse du serveur de destination
	Address string

	// description du lien
	Desc string
}

func (s *Server) FindTarget(address string) (Link, error) {
	for _, t := range s.Links {
		if t.Address == address {
			return t, nil
		}
	}
	return Link{}, errInvalidArgument
}

// Entry est une entrée dans une base de données
type Entry struct {
	// clef unique
	ID string

	// mots-clefs utilisés pour la recherche
	Keywords []string

	// niveau de privilège requis
	Restricted int

	// accessible uniquement au propriétaire
	Owner string

	// titre de l'entrée
	Title string

	// contenu de l'entrée
	Content string
}

func (s *Server) DataSearch(keyword string, owner string) []Entry {
	result := make([]Entry, 0, len(s.Entries))
	for _, e := range s.Entries {
		if e.Match(keyword) {
			if e.Owner == "" || e.Owner == owner {
				result = append(result, e)
			}
		}
	}
	return result
}

func (s *Server) FindEntry(id string, owner string) (Entry, error) {
	for _, e := range s.Entries {
		if e.ID == id {
			if e.Owner == "" || e.Owner == owner {
				return e, nil
			}
		}
	}
	return Entry{}, errInvalidArgument
}

// Match détermine si l'entrée contient le mot-clef
func (e Entry) Match(keyword string) bool {
	match := fuzzy.FindNormalizedFold(keyword, e.Keywords)
	return len(match) > 0
}

// Register représente registre mémoire qui peut être modifié pour contrôler quelque chose
type Register struct {
	Name        string
	State       bool
	Description string
	Restricted  int
}

func (r *Register) Match(name string) bool {
	return fuzzy.MatchNormalizedFold(name, r.Name)
}

func (s *Server) RegistrySearch(name string) []Register {
	result := make([]Register, 0, len(s.Registers))
	for _, r := range s.Registers {
		if r.Match(name) {
			result = append(result, r)
		}
	}
	return result
}

func normalize(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	out, _, _ := transform.String(t, s)
	return out
}

func (s *Server) RegistryEdit(name string) (bool, error) {
	for i, r := range s.Registers {
		if r.Name == name {
			s.Registers[i].State = !s.Registers[i].State
			return s.Registers[i].State, nil
		}
	}
	return false, fmt.Errorf("%s : %w", name, errRegisterNotFound)
}

// CreateBackdoor créé une backdoor dans le serveur
func (s *Server) CreateBackdoor(login string) {
	acc := Account{
		Login:    login,
		Admin:    false,
		Backdoor: true,
	}
	s.Accounts = append(s.Accounts, acc)
}

// Deal with Forum
func (s *Server) GetForum() (Forum, error) {
	return GetForum("toile/" + s.Address)
}
