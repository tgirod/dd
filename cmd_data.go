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
			Path:      []string{"data"},
			ShortHelp: "effectue une recherche par mot clef",
			Args: []Arg{
				{
					Name:      "keyword",
					ShortHelp: "mot clef utilisé pour la recherche",
					Type:      TextArg,
				},
			},
			Run: DataSearch,
		},
		{
			Name:      "view",
			Path:      []string{"data"},
			ShortHelp: "affiche le contenu d'une entrée",
			Args: []Arg{
				{
					Name:      "id",
					ShortHelp: "identifiant de l'entrée à afficher",
					Type:      TextArg,
				},
			},
			Run: DataView,
		},
	},
}

func DataSearch(ctx Context) any {
	result := ctx.Result()

	keyword := ctx.Args[0]
	if len([]rune(keyword)) < 3 {
		result.Error = fmt.Errorf("%s : %w", keyword, errKeywordTooShort)
		return result
	}

	// construire la réponse à afficher
	entries := ctx.Server.DataSearch(keyword, ctx.Identity.Login)
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
	result := ctx.Result()

	id := ctx.Args[0]
	entry, err := ctx.FindEntry(id, ctx.Identity.Login)
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