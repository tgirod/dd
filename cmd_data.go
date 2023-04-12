package main

type DataSearchMsg struct {
	Keyword string
}

type DataViewMsg struct {
	Id string
}

var data = Cmd{
	Name:      "data",
	ShortHelp: "recherche des données sur le serveur",
	Connected: true,
	SubCmds: []Cmd{
		{
			Name:      "search",
			Path:      []string{"data"},
			ShortHelp: "effectue une recherche par mot clef",
			Args: []Arg{
				{
					Name:      "keyword",
					ShortHelp: "mot clef utilisé pour la recherche",
				},
			},
			Run: func(ctx Context, args []string) any {
				return DataSearchMsg{args[0]}
			},
		},
		{
			Name:      "view",
			Path:      []string{"data"},
			ShortHelp: "affiche le contenu d'une entrée",
			Args: []Arg{
				{
					Name:      "id",
					ShortHelp: "identifiant de l'entrée à afficher",
				},
			},
			Run: func(ctx Context, args []string) any {
				return DataViewMsg{args[0]}
			},
		},
	},
}
