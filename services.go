package main

import (
	"unicode"

	//"github.com/lithammer/fuzzysearch/fuzzy"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Service regroupe les infos de base exposées par tous les services
type Service struct {
	// nom du service
	Name string

	// description courte du service pour l'index
	Description string

	// niveau de privilège requis pour utiliser le service
	Restricted int
}

// Gate est un service permettant de se connecter ailleurs
type Gate struct {
	// description du service
	Description string

	// liens fournis par le service
	Targets []Target
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

func (g Gate) IsEmpty() bool {
	return len(g.Targets) == 0
}

// Database est un service de base de données
type Database struct {
	// description du service
	Description string

	// données contenues
	Entries []Entry
}

func (d Database) IsEmpty() bool {
	return len(d.Entries) == 0
}

// Entry est une entrée dans une base de données
type Entry struct {
	// clef unique
	Key string

	// mots-clefs utilisés pour la recherche
	Keywords []string

	// niveau de privilège requis
	Restricted int

	// titre de l'entrée
	Title string

	// contenu de l'entrée
	Content string
}

// Search retourne la liste des entrées contenant le mot-clef
func (d Database) Search(keyword string, privilege int) []Entry {
	result := make([]Entry, 0, len(d.Entries))
	for _, e := range d.Entries {
		if e.Match(keyword) && privilege >= e.Restricted {
			result = append(result, e)
		}
	}
	return result
}

// Match détermine si l'entrée contient le mot-clef
func (e Entry) Match(keyword string) bool {
	keyword = normalize(keyword)

	for _, k := range e.Keywords {
		if k == keyword {
			return true
		}
	}

	return false
}

func normalize(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	out, _, _ := transform.String(t, s)
	return out
}
