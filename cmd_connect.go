package main

import (
	"fmt"
	"strings"
)

type ConnectMsg struct {
	Address string
}

var connect = Cmd{
	Name:       "connect",
	ShortHelp:  "établit la connexion avec un serveur",
	Args:       []Arg{{Name: "address", ShortHelp: "adresse sur serveur auquel se connecter", Type: ShortArg}},
	SubCmds:    []Cmd{},
	Connected:  false,
	Identified: false,
	Run:        Connect,
}

func Connect(ctx Context) any {
	console := ctx.Value("console").(*Console)
	address := ctx.Value("address").(string)
	result := ctx.Result()

	if err := console.Connect(address, false); err != nil {
		result.Error = err
		return result
	}

	console.History.Clear()
	console.History.Push(Link{address, ""})

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", console.Server.Address)
	fmt.Fprintf(&b, "%s\n", console.Server.Description)
	result.Output = b.String()
	return result
}