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
	console := ctx.Value("console").(*Console)
	address := ctx.Value("address").(string)

	if err := console.Connect(address, false); err != nil {
		return ctx.Result(err, "")
	}

	console.History.Clear()
	console.History.Push(Link{address, ""})

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", console.Server.Address)
	fmt.Fprintf(&b, "%s\n", console.Server.Description)
	return ctx.Result(nil, b.String())
}