package main

import (
	"fmt"
	"strings"
)

var data = Cmd{
	name:      "data",
	help:      "recherche des données sur le serveur",
	connected: true,
	next: Branch{
		name: "action",
		cmds: []Cmd{
			{
				name: "search",
				help: "effectue une recherche par mot clef",
				next: String{
					name: "keyword",
					help: "mot clef utilisé pour la recherche",
					next: Run(DataSearch),
				},
			},
			{
				name: "view",
				help: "affiche le contenu d'une entrée",
				next: String{
					name: "id",
					help: "identifiant de l'entrée à afficher",
					next: Run(DataView),
				},
			},
		},
	},
}

func DataSearch(ctx Context) any {
	console := ctx.Value("console").(*Console)
	keyword := ctx.Value("keyword").(string)

	if len([]rune(keyword)) < 3 {
		return ctx.Result(fmt.Errorf("%s : %w", keyword, errKeywordTooShort), "")
	}

	// construire la réponse à afficher
	entries := console.Server.DataSearch(keyword, console.Account)
	b := strings.Builder{}
	tw := tw(&b)
	fmt.Fprintf(tw, "ID\tKEYWORDS\tTITLE\t\n")
	for _, e := range entries {
		title := e.Title
		fmt.Fprintf(tw, "%s\t%s\t%s\t\n",
			e.ID,
			strings.Join(e.Keywords, " "),
			title,
		)
	}
	tw.Flush()

	return ctx.Result(nil, b.String())
}

func DataView(ctx Context) any {
	console := ctx.Value("console").(*Console)
	id := ctx.Value("id").(string)

	entry, err := console.FindEntry(id, console.Account)
	if err != nil {
		return ctx.Result(err, "")
	}

	// construire la réponse à afficher
	b := strings.Builder{}
	fmt.Fprintf(&b, "TITLE: %s\n", entry.Title)
	fmt.Fprintf(&b, "KEYWORDS: %s\n", strings.Join(entry.Keywords, " "))
	fmt.Fprintf(&b, "-------------------------------------\n")
	fmt.Fprintf(&b, entry.Content)

	return ctx.Result(nil, b.String())
}