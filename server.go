package main

import (
	"time"
	"unicode"

	"github.com/lithammer/fuzzysearch/fuzzy"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Server représente un serveur sur le Net
type Server struct {
	// Addresse du serveur sur le réseau
	Address string

	// liste de codes d'accès valides pour se connecter au serveur
	Credentials []Cred

	// informations affichées lors de la connexion
	Description string

	// durée du scan avant de se faire repérer par le serveur
	Scan time.Duration

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

	// identité utilisée pour se connecter
	Login    string
	Password string
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

func (s *Server) RegisterSearch(name string) []Register {
	result := make([]Register, 0, len(s.Registers))
	for _, r := range s.Registers {
		if r.Match(name) {
			result = append(result, r)
		}
	}
	return result
}

func (s *Server) FindRegister(name string) (*Register, error) {
	for i, r := range s.Registers {
		if r.Name == name {
			return &s.Registers[i], nil
		}
	}
	return nil, errInvalidArgument
}

func normalize(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	out, _, _ := transform.String(t, s)
	return out
}

// Deal with Forum
func (s *Server) GetForum() (Forum, error) {
	return GetForum( "toile/"+s.Address );
}
