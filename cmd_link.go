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
		name: "id",
		help: "identifiant du lien",
		options: func(ctx Context) []Option {
			console := ctx.Value("console").(*Console)
			links := console.Server.Links
			opts := make([]Option, len(links))
			for i, l := range links {
				opts[i].desc = l.Desc
				opts[i].value = i
			}
			return opts
		},
		next: Run(LinkCmd),
	},
}

func LinkCmd(ctx Context) any {
	console := ctx.Value("console").(*Console)
	id := ctx.Value("id").(int)

	if id < 0 || id >= len(console.Server.Links) {
		return ctx.Result(
			fmt.Errorf("%v : %w", id, errInvalidArgument),
			"",
		)
	}

	link := console.Server.Links[id]

	if err := console.Connect(link.Address, false); err != nil {
		return ctx.Result(
			err,
			"",
		)
	}

	console.History.Push(link)

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", console.Server.Address)
	fmt.Fprintf(&b, "%s\n", console.Server.Description)
	return ctx.Result(
		nil,
		b.String(),
	)
}
