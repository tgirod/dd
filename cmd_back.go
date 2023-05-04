package main

import (
	"fmt"
	"strings"
)

var back = Cmd{
	name:       "back",
	help:       "quitte le serveur actuel et se reconnecte au serveur précédent",
	connected:  true,
	identified: false,
	next:       Run(Back),
}

func Back(ctx Context) any {
	console := ctx.Value("console").(*Console)
	if console.Session == nil || console.Session.Parent == nil {
		return ctx.Result(errInvalidCommand, "")
	}

	// enlever le serveur actuel
	console.Session = console.Session.Parent

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", console.Address)
	return ctx.Result(nil, b.String())
}
