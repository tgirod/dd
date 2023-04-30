package main

import (
	"fmt"
	"strings"
)

var jack = Cmd{
	Name:      "jack",
	ShortHelp: "force l'accès à un lien",
	Connected: true,
	Args:      link.Args,
	Run:       Jack,
}

func Jack(ctx Context) any {
	console := ctx.Value("console").(*Console)
	res := ctx.Result()

	id := ctx.Value("id").(int)

	if id < 0 || id >= len(console.Server.Links) {
		res.Error = errInvalidArgument
		return res
	}

	link := console.Server.Links[id]

	if err := console.Connect(link.Address, true); err != nil {
		res.Error = err
		return res
	}

	console.History.Push(link)
	console.StartAlert()

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", console.Server.Address)
	fmt.Fprintf(&b, "%s\n", console.Server.Description)
	res.Output = b.String()
	return res
}