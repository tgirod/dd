package main

import (
	"fmt"
	"strings"
)

type IndexMsg struct{}

var index = Cmd{
	Name:      "index",
	ShortHelp: "liste les services disponibles dans le serveur courant",
	Connected: true,
	Run:       Index,
}

func Index(ctx Context) any {
	console := ctx.Value("console").(*Console)
	res := ctx.Result()

	b := strings.Builder{}

	s := console.Server
	b.WriteString(s.Description)
	b.WriteString("\n")
	fmt.Fprintf(&b, "LIENS     : %d\n", len(s.Links))
	fmt.Fprintf(&b, "DONNEES   : %d\n", len(s.Entries))
	fmt.Fprintf(&b, "REGISTRES : %d\n", len(s.Registers))

	res.Output = b.String()
	return res
}