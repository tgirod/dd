package main

import (
	"fmt"
	"strings"
)

var link = Cmd{
	name:       "link",
	help:       "utilise les liens pour se connecter à un autre serveur",
	connected:  true,
	identified: false,
	next: Select{
		name:   "id",
		help:   "identifiant du lien",
		header: "liste des liens disponibles dans ce serveur",
		options: func(ctx Context) []Option {
			console := ctx.Value("console").(*Console)
			links := console.Server.Links(console.Account)
			opts := make([]Option, len(links))
			for i, l := range links {
				opts[i].help = l.Desc
				opts[i].value = l.ID
			}
			return opts
		},
		next: Run(LinkCmd),
	},
}

func LinkCmd(ctx Context) any {
	console := ctx.Value("console").(*Console)
	id := ctx.Value("id").(int)

	link, err := console.Server.Link(id, console.Account)
	if err != nil {
		return ctx.Error(err)
	}

	if err := console.Connect(link.Address, false); err != nil {
		return ctx.Error(err)
	}

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", console.Server.Address)
	fmt.Fprintf(&b, "%s\n", console.Server.Description)
	return ctx.Output(b.String())
}
