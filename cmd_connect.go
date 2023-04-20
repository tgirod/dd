package main

import (
	"fmt"
	"strings"
)

type ConnectMsg struct {
	Address string
}

var connect = Cmd{
	Name:      "connect",
	ShortHelp: "établit la connexion avec un serveur",
	Args: []Arg{
		{
			Name:      "address",
			ShortHelp: "adresse sur serveur auquel se connecter",
			Type:      Text,
		},
	},
	Run: Connect,
}

func Connect(ctx Context) any {
	result := ctx.Result()
	args := ctx.Args

	address := args[0]
	if err := ctx.Console.Connect(address, false); err != nil {
		result.Error = err
		return result
	}

	ctx.History.Clear()
	ctx.History.Push(Link{address, ""})

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", ctx.Server.Address)
	fmt.Fprintf(&b, "%s\n", ctx.Server.Description)
	result.Output = b.String()
	return result
}