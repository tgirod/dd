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
	console := ctx.Console()
	s := console.Server

	b := strings.Builder{}
	tw := tw(&b)

	// récupérer les liens
	links, err := Find[Link](
		s.Match(),
	)
	if err != nil {
		return ctx.Error(err)
	}

	// récupérer les registres
	registers, err := Find[Register](
		s.Match(),
	)
	if err != nil {
		return ctx.Error(err)
	}

	fmt.Fprintf(tw, "LINKS\n")
	fmt.Fprintf(tw, "ID\tGroup\tAddress\tDescription\t\n")
	for _, l := range links {
		fmt.Fprintf(tw, "%d\t%s\t%s\t%s\t\n", l.ID, l.Group, l.Address, l.Desc)
	}
	fmt.Fprintln(tw)

	fmt.Fprintf(tw, "REGISTERS\n")
	fmt.Fprintf(tw, "ID\tGroup\tDescription\tState\t\n")
	for _, r := range registers {
		fmt.Fprintf(tw, "%d\t%s\t%s\t%s\t\n", r.ID, r.Group, r.Description, r.State)
	}

	tw.Flush()
	return ctx.Output(b.String())
}
