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
		name:   "id",
		help:   "identifiant du lien",
		header: "liste des liens disponibles dans ce serveur",
		options: func(ctx Context) ([]Option, error) {
			console := ctx.Console()
			links := console.Server.Links(console.User)
			return ToOptions(links), nil
		},
		next: Run(func(ctx Context) any {
			ctx = ctx.WithContext(Run(Jack), "", nil)
			m := NewLoader(ctx, 5*time.Second, []string{"hack en cours"})
			return OpenModalMsg(m)
		}),
	},
}

func Jack(ctx Context) any {
	console := ctx.Console()

	id := ctx.Value("id").(int)

	link, err := console.Server.Link(id, console.User)
	if err != nil {
		return ctx.Error(err)
	}

	if err := console.Connect(link.Address, console.Identity, true); err != nil {
		return ctx.Error(err)
	}

	console.StartAlert()

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", console.Server.Address)
	fmt.Fprintf(&b, "%s\n", console.Server.Description)
	return ctx.Output(b.String())
}
