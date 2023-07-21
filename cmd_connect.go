package main

import (
	"fmt"
	"strings"
)

type ConnectMsg struct {
	Address string
}

var connect = Cmd{
	name:       "connect",
	help:       "établit la connexion avec un serveur",
	connected:  false,
	identified: false,
	next: String{
		name: "address",
		help: "adresse du serveur auquel se connecter",
		next: Run(Connect),
	},
}

func Connect(ctx Context) any {
	console := ctx.Console()
	address := ctx.Value("address").(string)

	fmt.Printf("LOG try connecting\n")
	if err := console.Connect(address, console.Identity, false, true); err != nil {
		// connexion impossible avec l'identité courante
		// lancer la saisie du login
		fmt.Printf("LOG Connect bad identity\n")
		return ctx.WithContext(idconnect, "address", address)
	}

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", console.Server.Address)
	fmt.Fprintf(&b, "%s\n", console.Server.Description)
	return ctx.Result(nil, b.String())
}

var idconnect = Text{
	name: "login",
	help: "identifiant utilisateur",
	next: Hidden{
		name: "password",
		help: "mot de passe utilisateur",
		next: Run(IdConnect),
	},
}

func IdConnect(ctx Context) any {
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
	err = console.Connect(address, identity, false, true)
	if err != nil {
		return ctx.Error(err)
	}

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", console.Server.Address)
	fmt.Fprintf(&b, "%s\n", console.Server.Description)
	return ctx.Result(nil, b.String())
}
