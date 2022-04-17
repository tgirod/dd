package main

import "github.com/lithammer/fuzzysearch/fuzzy"

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
	Service

	// adresse du serveur distant
	TargetAddress string

	// niveau de privilège obtenu une fois connecté
	Privilege int
}

// Database est un service de base de données
type Database struct {
	// description du service
	Service

	// données contenues
	Entries []Entry
}

// Entry est une entrée dans une base de données
type Entry struct {
	// mots-clefs utilisés pour la recherche
	Keywords []string

	// titre de l'entrée
	Title string

	// contenu de l'entrée
	Content string
}

// Match détermine si l'entrée contient le mot-clef
func (e Entry) Match(keyword string) bool {
	find := fuzzy.FindNormalizedFold(keyword, e.Keywords)
	return len(find) > 0
}

// Search retourne la liste des entrées contenant le mot-clef
func (d Database) Search(keyword string) []Entry {
	result := make([]Entry, 0, len(d.Entries))
	for _, e := range d.Entries {
		if e.Match(keyword) {
			result = append(result, e)
		}
	}
	return result
}
