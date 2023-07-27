package main

import (
	"errors"
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
		options: func(ctx Context) ([]Option, error) {
			console := ctx.Console()
			links := console.Server.Links(console.User)
			return ToOptions(links), nil
		},
		next: Run(LinkCmd),
	},
}

func LinkCmd(ctx Context) any {
	console := ctx.Console()
	id := ctx.Value("id").(int)

	link, err := console.Server.Link(id, console.User)
	if err != nil {
		return ctx.Error(err)
	}

	if err := console.Connect(link.Address, console.Identity, false, false); err != nil {
		if errors.Is(err, errInvalidUser) {
			// connexion impossible avec l'identité courante
			return ctx.WithContext(idlink, "address", link.Address)
		}

		return ctx.Error(err)
	}

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", console.Server.Address)
	fmt.Fprintf(&b, "%s\n", console.Server.Description)
	return ctx.Output(b.String())
}

var idlink = String{
	name: "address",
	help: "addresse du serveur",
	next: Text{
		name: "login",
		help: "identifiant utilisateur",
		next: Hidden{
			name: "password",
			help: "mot de passe utilisateur",
			next: Run(IdLink),
		},
	},
}

func IdLink(ctx Context) any {
	console := ctx.Console()

	// vérifier la validité de l'identité
	login := ctx.Value("login").(string)
	password := ctx.Value("password").(string)
	identity, err := CheckIdentity(login, password)
	if err != nil {
		return ctx.Error(err)
	}

	// effectuer la connexion
	address := ctx.Value("address").(string)
	err = console.Connect(address, identity, false, false)
	if err != nil {
		return ctx.Error(err)
	}

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", console.Server.Address)
	fmt.Fprintf(&b, "%s\n", console.Server.Description)
	return ctx.Result(nil, b.String())
}
