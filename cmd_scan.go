package main

import (
	"fmt"
	"strings"
)

type ScanMsg struct{}

var scan = Cmd{
	name:      "scan",
	help:      "affiche des informations sur les ressources innaccessibles dans le serveur",
	connected: true,
	next:      Run(Scan),
}

func Scan(ctx Context) any {
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
