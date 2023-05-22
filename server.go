package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/lithammer/fuzzysearch/fuzzy"

	"gopkg.in/yaml.v3"
)

// Server représente un serveur sur le Net
type Server struct {
	// Addresse du serveur sur le réseau
	Address string `storm:"id"`

	// ce serveur accepte-t-il des connexions anonymes ?
	Private bool

	// informations affichées lors de la connexion
	Description string

	// temps que met le serveur à effectuer la trace
	Security time.Duration
}

func (s Server) Match() q.Matcher {
	return q.Eq("Server", s.Address)
}

// User représente un compte utilisateur sur un serveur
type User struct {
	ID       int    `storm:"id,increment"`
	Login    string `storm:"index"`
	Server   string `storm:"index"` // le serveur concerné
	Backdoor bool
	Groups   []string `storm:"index"`
}

func (u User) Match() q.Matcher {
	return q.NewFieldMatcher("Group", u)
}

// Match permet de vérifier si une donnée est accessible depuis un compte
func (u User) MatchField(v any) (bool, error) {
	group, ok := v.(string)
	if !ok {
		return false, storm.ErrBadType
	}

	// aucun groupe == public
	if group == "" {
		return true, nil
	}

	for _, g := range u.Groups {
		if group == g {
			return true, nil
		}
	}
	return false, nil
}

func (s Server) Users() []User {
	users, err := Find[User](s.Match())
	if err != nil {
		panic(err)
	}
	return users
}

// FindUser cherche un compte utilisateur correspondant au login
func (s Server) FindUser(login string) (User, error) {
	return First[User](
		s.Match(),
		q.Eq("Login", login),
	)
}

func (s Server) RemoveUser(user User) error {
	return Delete(user)
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

func (s Server) Links(a User) []Link {
	links, err := Find[Link](
		s.Match(),
		a.Match(),
	)
	if err != nil {
		panic(err)
	}
	return links
}

func (s Server) Link(id int, a User) (Link, error) {
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

func (s Server) Entries(a User) []Entry {
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

func (s Server) DataSearch(keyword string, a User) []Entry {
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

func (s Server) FindEntry(id string, a User) (Entry, error) {
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

func (s Server) Registers(a User) []Register {
	registers, err := Find[Register](
		s.Match(),
		a.Match(),
	)
	if err != nil {
		panic(err)
	}
	return registers
}

func (s Server) Register(id int, a User) (Register, error) {
	return First[Register](
		s.Match(),
		a.Match(),
		q.Eq("ID", id),
	)
}

// CreateBackdoor créé une backdoor dans le serveur
func (s Server) CreateBackdoor(identity Identity) (User, error) {
	acc := User{
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
func (p Post) Dump() {
	fmt.Printf("--- Dump Post:")
	fmt.Printf("\n Server: [%s]", p.Server)
	fmt.Printf("\n Group: [%s]", p.Group)
	fmt.Printf("\n ID: [%d]", p.ID)
	fmt.Printf("\n Parent: [%d]", p.Parent)
	fmt.Printf("\n Date: [%s]", p.Date)
	fmt.Printf("\n Author: [%s]", p.Author)
	fmt.Printf("\n Subject: [%s]", p.Subject)
	fmt.Printf("\n Content: [%v]", p.Content)
}

// func (s Server) Posts(a User) []Post {
// 	posts, err := Find[Post](
// 		s.Match(),
// 		a.Match(),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return posts
// }

// TEST : serialize all Posts to YAML
func SerializePosts(addr string) {
	s, err := FindServer(addr)
	if err != nil {
		panic(err)
	}

	posts, err := Find[Post](
		s.Match(),
	)
	if err != nil {
		panic(err)
	}

	// all posts
	d, err := yaml.Marshal(posts)
		if err != nil {
		panic(err)
	}
	fmt.Printf("--- all posts:\n%s\n\n", d)
}
// TEST Load new post from YAML file
func LoadPosts(path string) {
	buf, err := ioutil.ReadFile(path)
		if err != nil {
		panic(err)
	}

	p := Post{}
	p.Dump()
	fmt.Printf("--- New Post:\n%v\n", p)

	err = yaml.Unmarshal(buf, &p)
		if err != nil {
		panic(err)
	}
	fmt.Print("** Unmarshal\n")
	p.Dump()

	post, err := Save(p)
	if err != nil {
		panic(err)
	}

	fmt.Print("** Saving\n")
	post.Dump()
}

//func (s Server) Post(id int, a User) (Post, error) {
func (s Server) Post(id int) (Post, error) {
	return First[Post](
		s.Match(),
		//a.Match(),
		q.Eq("ID", id),
	)
}

// Topics liste les posts qui n'ont pas de parent
// func (s Server) Topics(a User) []Post {
func (s Server) Topics() []Post {
	posts, err := Find[Post](
		s.Match(),
		// a.Match(),
		q.Eq("Parent", 0),
	)
	if err != nil {
		panic(err)
	}
	return posts
}

// Replies retourne la liste des réponses à un post
// func (s Server) Replies(parent int, a User) []Post {
func (s Server) Replies(parent int) []Post {
	posts, err := Find[Post](
		s.Match(),
		// a.Match(),
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

// func (s Server) RecReplies(parent int, a User) []Post {
// 	thread, err := Find[Post](
// 		s.Match(),
// 		a.Match(),
// 		q.Eq("Parent", parent),
// 	)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// insérer les réponses
// 	for i, p := range thread {
// 		rec := s.RecReplies(p.ID, a)
// 		thread = concat(
// 			thread[:i+1],
// 			rec,
// 			thread[i+1:],
// 		)
// 	}

// 	return thread
// }

// type Thread struct {
// 	Post
// 	Replies []Thread
// }
type Thread struct {
	Post
	Replies []Post
}
func (s Server) Thread(p Post) (Thread, error) {

	replies, err := Find[Post](
		s.Match(),
		q.Eq("Parent", p.ID),
	)

	thread := Thread{Post: p}
	if err != nil {
		return thread, err
	}

	thread.Replies = append(thread.Replies, replies...)
	return thread, nil
}

// func (s Server) Thread(p Post, a User) (Thread, error) {
// 	replies, err := Find[Post](
// 		s.Match(),
// 		a.Match(),
// 		q.Eq("Parent", p.ID),
// 	)

// 	thread := Thread{Post: p}
// 	if err != nil {
// 		return thread, err
// 	}

// 	for _, r := range replies {
// 		sub, err := s.Thread(r, a)
// 		if err != nil {
// 			return thread, err
// 		}
// 		thread.Replies = append(thread.Replies, sub)
// 	}

// 	return thread, nil
// }
