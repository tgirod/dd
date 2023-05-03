package main

import (
	"fmt"
	"strings"
	"time"
)

var jack = Cmd{
	name:      "jack",
	help:      "force l'accès à un lien",
	connected: true,
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
		next: Run(func(ctx Context) any {
			ctx = ctx.WithContext(Run(Jack), "", nil)
			m := NewLoader(ctx, 5*time.Second, []string{"hack en cours"})
			return OpenModalMsg(m)
		}),
	},
}

func Jack(ctx Context) any {
	console := ctx.Value("console").(*Console)

	id := ctx.Value("id").(int)

	link := console.Server.Links[id]

	if err := console.Connect(link.Address, true); err != nil {
		return ctx.Error(err)
	}

	console.History.Push(link)
	console.StartAlert()

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", console.Server.Address)
	fmt.Fprintf(&b, "%s\n", console.Server.Description)
	return ctx.Output(b.String())
}