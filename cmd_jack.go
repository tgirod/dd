package main

import (
	"fmt"
	"strings"
)

var jack = Cmd{
	name: "jack",
	help: "force l'accès à un serveur",
	next: Branch{
		name: "",
		cmds: []Cmd{
			{
				name:      "link",
				help:      "force l'accès via un lien",
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
					next: Run(JackLink),
				},
			},
			{
				name: "connect",
				help: "force l'accès via une adresse",
				next: String{
					name: "address",
					help: "adresse du serveur auquel se connecter",
					next: Run(JackConnect),
				},
			},
		},
	},
}

func JackLink(ctx Context) any {
	console := ctx.Console()

	id := ctx.Value("id").(int)

	link, err := console.Server.Link(id, console.User)
	if err != nil {
		return ctx.Error(err)
	}

	if err := console.Connect(link.Address, console.Identity, true, false); err != nil {
		return ctx.Error(err)
	}

	console.StartAlert()

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", console.Server.Address)
	fmt.Fprintf(&b, "%s\n", console.Server.Description)
	return ctx.Output(b.String())
}

func JackConnect(ctx Context) any {
	console := ctx.Console()
	address := ctx.Value("address").(string)

	if err := console.Connect(address, console.Identity, true, true); err != nil {
		return ctx.Error(err)
	}

	console.StartAlert()

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", console.Server.Address)
	fmt.Fprintf(&b, "%s\n", console.Server.Description)
	return ctx.Result(nil, b.String())
}
