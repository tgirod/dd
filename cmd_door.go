package main

import (
	"fmt"
	"strings"
)

type DoorMsg struct{}

var door = Cmd{
	name:      "door",
	help:      "créé une backdoor dans le serveur",
	connected: true,
	next:      Run(Door),
}

func Door(ctx Context) any {
	console := ctx.Value("console").(*Console)

	// créer une nouvelle identité aléatoire
	id := console.CreateRandomIdentity()

	// créer une backdoor associée dans le serveur
	console.CreateBackdoor(id.Login)

	b := strings.Builder{}
	fmt.Fprintf(&b, "backdoor créée sur le serveur %s\n", console.Server.Address)
	fmt.Fprintf(&b, "login: %s\n", id.Login)
	fmt.Fprintf(&b, "password: %s\n", id.Password)
	fmt.Fprintf(&b, "cette backdoor sera détruite automatiquement après usage.\n")

	return ctx.Result(nil, b.String())
}