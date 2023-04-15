package main

import (
	"fmt"
	"strings"
)

type DoorMsg struct{}

var door = Cmd{
	Name:      "door",
	ShortHelp: "créé une backdoor dans le serveur",
	Connected: true,
	Run:       Door,
}

func Door(ctx Context) any {
	result := ctx.Result()

	// créer une nouvelle identité aléatoire
	id := ctx.CreateRandomIdentity()

	// créer une backdoor associée dans le serveur
	ctx.CreateBackdoor(id.Login)

	b := strings.Builder{}
	fmt.Fprintf(&b, "backdoor créée sur le serveur %s\n", ctx.Server.Address)
	fmt.Fprintf(&b, "login: %s\n", id.Login)
	fmt.Fprintf(&b, "password: %s\n", id.Password)
	fmt.Fprintf(&b, "cette backdoor sera détruite automatiquement après usage.\n")

	result.Output = b.String()
	return result
}