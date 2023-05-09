package main

import (
	"fmt"
	"strings"
)

// du code partagé par link et connect pour pouvoir saisir un login/password
// quand on tente de se connecter à un serveur privé

var idconnect = Cmd{
	name: "idconnect",
	help: "",
	next: Text{
		name: "login",
		help: "identifiant utilisateur",
		next: Hidden{name: "password",
			help: "mot de passe utilisateur",
			next: Run(IdConnect),
		},
	},
}

func IdConnect(ctx Context) any {
	console := ctx.Console()
	address := ctx.Value("address").(string)
	login := ctx.Value("login").(string)
	password := ctx.Value("password").(string)

	identity, err := CheckIdentity(login, password)
	if err != nil {
		return ctx.Error(err)
	}

	err = console.Connect(address, identity, false)
	if err != nil {
		return ctx.Error(err)
	}

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", console.Server.Address)
	fmt.Fprintf(&b, "%s\n", console.Server.Description)
	return ctx.Result(nil, b.String())
}
