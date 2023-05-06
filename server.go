package main

import (
	"errors"
	"time"

	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

// Server représente un serveur sur le Net
type Server struct {
	// Addresse du serveur sur le réseau
	Address string `storm:"id"`

	// ce serveur accepte-t-il des connexions anonymes ?
	Public bool

	// informations affichées lors de la connexion
	Description string

	// durée du scan avant de se faire repérer par le serveur
	Scan time.Duration
}

func (s Server) Match() q.Matcher {
	return q.Eq("Server", s.Address)
}

// Account représente un compte utilisateur sur un serveur
type Account struct {
	Login    string `storm:"id"`
	Server   string `storm:"index"` // le serveur concerné
	Backdoor bool
	Groups   []string `storm:"index"`
}

func (a Account) Match() q.Matcher {
	return q.NewFieldMatcher("Group", a)
}

// Match permet de vérifier si une donnée est accessible depuis un compte
func (a Account) MatchField(v any) (bool, error) {
	group, ok := v.(string)
	if !ok {
		return false, storm.ErrBadType
	}

	// aucun groupe == public
	if group == "" {
		return true, nil
	}

	for _, g := range a.Groups {
		if group == g {
			return true, nil
		}
	}
	return false, nil
}

func (s Server) Accounts() []Account {
	accounts, err := Find[Account](s.Match())
	if err != nil {
		panic(err)
	}
	return accounts
}

// FindAccount cherche un compte utilisateur correspondant au login
func (s Server) FindAccount(login string) (Account, error) {
	return First[Account](
		s.Match(),
		q.Eq("Login", login),
	)
}

func (s Server) RemoveAccount(account Account) error {
	return Delete(account)
}

type Link struct {
	ID     int    `storm:"id,increment"`
	Server string `storm:"index"`
	Group  string `storm:"index"`

	// adresse du serveur de destination
	Address string

	// description du lien
	Desc string
}

func (s Server) Links(a Account) []Link {
	links, err := Find[Link](
		s.Match(),
		a.Match(),
	)
	if err != nil {
		panic(err)
	}
	return links
}

func (s Server) Link(id int, a Account) (Link, error) {
	return First[Link](
		s.Match(),
		a.Match(),
		q.Eq("ID", id),
	)
}

// Entry est une entrée dans une base de données
type Entry struct {
	Server string `storm:"index"`
	Group  string `storm:"index"`

	// identifiant unique
	ID string `storm:"id"`

	// mots-clefs utilisés pour la recherche
	Keywords []string `storm:"index"`

	// accessible uniquement au propriétaire
	Owner string `storm:"index"`

	// titre de l'entrée
	Title string

	// contenu de l'entrée
	Content string
}

func (s Server) Entries(a Account) []Entry {
	entries, err := Find[Entry](
		s.Match(),
		a.Match(),
	)
	if err != nil {
		panic(err)
	}
	return entries
}

type KeywordMatcher string

func (m KeywordMatcher) Match(v any) (bool, error) {
	entry, ok := v.(Entry)
	if !ok {
		return false, errors.New("type incompatible")
	}
	return entry.Match(string(m)), nil
}

func (s Server) DataSearch(keyword string, a Account) []Entry {
	entries, err := Find[Entry](
		s.Match(),
		a.Match(),
		KeywordMatcher(keyword),
	)

	if err != nil {
		panic(err)
	}

	return entries
}

func (s Server) FindEntry(id string, a Account) (Entry, error) {
	return First[Entry](
		s.Match(),
		a.Match(),
		q.Eq("ID", id),
	)
}

// Match détermine si l'entrée contient le mot-clef
func (e Entry) Match(keyword string) bool {
	match := fuzzy.FindNormalizedFold(keyword, e.Keywords)
	return len(match) > 0
}

// Register représente registre mémoire qui peut être modifié pour contrôler quelque chose
type Register struct {
	Server      string `storm:"index"`
	Group       string `storm:"index"`
	ID          int    `storm:"id,increment"`
	Description string
	State       string   // état actuel
	Options     []string // valeurs possible
}

func (s Server) Registers(a Account) []Register {
	registers, err := Find[Register](
		s.Match(),
		a.Match(),
	)
	if err != nil {
		panic(err)
	}
	return registers
}

func (s Server) Register(id int, a Account) (Register, error) {
	return First[Register](
		s.Match(),
		a.Match(),
		q.Eq("ID", id),
	)
}

// CreateBackdoor créé une backdoor dans le serveur
func (s Server) CreateBackdoor(identity Identity) (Account, error) {
	acc := Account{
		Login:    identity.Login,
		Server:   s.Address,
		Backdoor: true,
	}
	return Save(acc)
}

type Post struct {
	Server  string `storm:"index"`
	Group   string `storm:"index"`
	ID      int    `storm:"id,increment"`
	Parent  int    `storm:"index"`
	Date    time.Time
	Author  string
	Subject string
	Content string
}

func (s Server) Posts(a Account) []Post {
	posts, err := Find[Post](
		s.Match(),
		a.Match(),
	)
	if err != nil {
		panic(err)
	}
	return posts
}

func (s Server) Post(id int, a Account) (Post, error) {
	return First[Post](
		s.Match(),
		a.Match(),
		q.Eq("ID", id),
	)
}

// Topics liste les posts qui n'ont pas de parent
func (s Server) Topics(a Account) []Post {
	posts, err := Find[Post](
		s.Match(),
		a.Match(),
		q.Eq("Parent", 0),
	)
	if err != nil {
		panic(err)
	}
	return posts
}

// Replies retourne la liste des réponses à un post
func (s Server) Replies(parent int, a Account) []Post {
	posts, err := Find[Post](
		s.Match(),
		a.Match(),
		q.Eq("Parent", parent),
	)
	if err != nil {
		panic(err)
	}

	return posts
}

func concat[T any](slices ...[]T) []T {
	var res []T
	for _, s := range slices {
		res = append(res, s...)
	}
	return res
}

func (s Server) Thread(parent int, a Account) []Post {
	thread, err := Find[Post](
		s.Match(),
		a.Match(),
		q.Eq("Parent", parent),
	)
	if err != nil {
		panic(err)
	}

	// insérer les réponses
	for i, p := range thread {
		rec := s.Thread(p.ID, a)
		thread = concat(
			thread[:i+1],
			rec,
			thread[i+1:],
		)
	}

	return thread
}
