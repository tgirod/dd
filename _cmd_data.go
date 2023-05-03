package main

import (
	"fmt"
	"strings"
)

var data = Cmd{
	Name:      "data",
	ShortHelp: "recherche des données sur le serveur",
	Connected: true,
	SubCmds: []Cmd{
		{
			Name:      "search",
			ShortHelp: "effectue une recherche par mot clef",
			Args: []Arg{
				{
					Name:      "keyword",
					ShortHelp: "mot clef utilisé pour la recherche",
					Type:      ShortArg,
				},
			},
			Run: DataSearch,
		},
		{
			Name:      "view",
			ShortHelp: "affiche le contenu d'une entrée",
			Args: []Arg{
				{
					Name:      "id",
					ShortHelp: "identifiant de l'entrée à afficher",
					Type:      ShortArg,
				},
			},
			Run: DataView,
		},
	},
}

func DataSearch(ctx Context) any {
	console := ctx.Value("console").(*Console)
	keyword := ctx.Value("keyword").(string)
	result := ctx.Result()

	if len([]rune(keyword)) < 3 {
		result.Error = fmt.Errorf("%s : %w", keyword, errKeywordTooShort)
		return result
	}

	// construire la réponse à afficher
	entries := console.Server.DataSearch(keyword, console.Identity.Login)
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

	result.Output = b.String()
	return result
}

func DataView(ctx Context) any {
	console := ctx.Value("console").(*Console)
	id := ctx.Value("id").(string)
	result := ctx.Result()

	entry, err := console.FindEntry(id, console.Identity.Login)
	if err != nil {
		result.Error = err
		return result
	}

	// construire la réponse à afficher
	b := strings.Builder{}
	fmt.Fprintf(&b, "TITLE: %s\n", entry.Title)
	fmt.Fprintf(&b, "KEYWORDS: %s\n", strings.Join(entry.Keywords, " "))
	fmt.Fprintf(&b, "-------------------------------------\n")
	fmt.Fprintf(&b, entry.Content)

	result.Output = b.String()
	return result
}