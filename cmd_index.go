package main

import (
	"fmt"
	"strings"
)

type IndexMsg struct{}

var index = Cmd{
	name:      "index",
	help:      "liste les services disponibles dans le serveur courant",
	connected: true,
	next:      Run(Index),
}

func Index(ctx Context) any {
	console := ctx.Value("console").(*Console)

	b := strings.Builder{}

	s := console.Server
	b.WriteString(s.Description)
	b.WriteString("\n")
	fmt.Fprintf(&b, "LIENS     : %d\n", len(s.Links()))
	fmt.Fprintf(&b, "DONNEES   : %d\n", len(s.Entries()))
	fmt.Fprintf(&b, "REGISTRES : %d\n", len(s.Registers()))

	return ctx.Result(nil, b.String())
}