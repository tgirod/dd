package main

import (
	"fmt"
	"math/rand"
	//"github.com/lithammer/fuzzysearch/fuzzy"
)

// Game contient l'état du jeu et les méthodes utiles pour en simplifier l'accès
type Game struct {
	Network []Server
}

func (g Game) FindServer(address string) (*Server, error) {
	for _, server := range g.Network {
		if server.Address == address {
			return &server, nil
		}
	}
	return nil, fmt.Errorf("%s : %w", address, errServerNotFound)
}

// Console représente le terminal depuis lequel le joueur accède au net
type Console struct {
	ID int

	// racine de l'arbre des commandes
	Node

	// identifiant dans le serveur actuel
	Login string

	// niveau de privilège dans le serveur actuel
	Privilege int

	// niveau d'alerte du serveur
	Alert int

	// zones mémoires disponibles pour une évasion
	Mem map[string]bool

	// interface neurale directe
	DNI bool

	// serveur auquel la console est connectée
	*Server
}

var Hack = map[string]Command{
	"jack":  Jack{},
	"rise":  Rise{},
	"evade": Evade{},
}

func NewConsole() *Console {
	// TODO compléter les commandes par défaut
	return &Console{
		Node: Node{
			Sub: []Command{
				Connect{},
				Node{
					Name: "data",
					Help: "effectuer une recherche sur le serveur",
					Sub: []Command{
						DataSearch{},
						DataView{},
					},
				},
				Help{},
				Index{},
				Link{},
				Load{},
				Plug{},
				Quit{},
			},
		},
	}
}

func (c *Console) IsConnected() bool {
	return c.Server != nil
}

// Illegal est une méthode appelée à chaque fois qu'une commande illégale est utilisée
func (c *Console) Illegal() {
	if len(c.Mem) == 0 {
		c.InitMem()
	}
}

func (c *Console) InitMem() {
	c.Mem = make(map[string]bool)
	for i := 0; i < 5; i++ {
		addr := fmt.Sprintf("%08x", rand.Uint32())
		c.Mem[addr] = true
	}
}

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

func (s *Server) Search(keyword string) []Entry {
	result := make([]Entry, 0, len(s.Entries))
	for _, e := range s.Entries {
		if e.Match(keyword) {
			result = append(result, e)
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
