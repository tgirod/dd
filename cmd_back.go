package main

import (
	"errors"
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

var errCannotBack = errors.New("vous êtes déjà sur le premier serveur")

func Back(ctx Context) any {
	console := ctx.Value("console").(*Console)

	// on ne peux pas reculer plus loin que le premier serveur
	parent := *console.Session.Parent
	if parent.Parent == nil {
		return ctx.Error(errCannotBack)
	}

	// enlever le serveur actuel
	console.Session = *console.Session.Parent

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", console.Address)
	return ctx.Result(nil, b.String())
}
