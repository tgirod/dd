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
	fmt.Fprintf(&b, "LIENS     : %d\n", len(s.Links(console.User)))
	fmt.Fprintf(&b, "DONNEES   : %d\n", len(s.Entries(console.User)))
	fmt.Fprintf(&b, "REGISTRES : %d\n", len(s.Registers(console.User)))

	return ctx.Result(nil, b.String())
}
