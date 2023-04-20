package main

import (
	"fmt"
	"strconv"
	"strings"
)

var jack = Cmd{
	Name:      "jack",
	ShortHelp: "force l'accès à un lien",
	Connected: true,
	Args: []Arg{
		{
			Name:      "id",
			ShortHelp: "identifiant du lien",
			Type:      LinkId,
		},
	},
	Run: Jack,
}

func Jack(ctx Context) any {
	res := ctx.Result()

	id, err := strconv.Atoi(ctx.Args[0])
	if err != nil {
		res.Error = errInvalidArgument
		return res
	}

	if id < 0 || id >= len(ctx.Server.Links) {
		res.Error = errInvalidArgument
		return res
	}

	address := ctx.Server.Links[id].Address
	if err := ctx.Console.Connect(address, true); err != nil {
		res.Error = err
		return res
	}

	ctx.History.Push(Link{address, ""})
	ctx.StartAlert()

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", ctx.Server.Address)
	fmt.Fprintf(&b, "%s\n", ctx.Server.Description)
	res.Output = b.String()
	return res
}