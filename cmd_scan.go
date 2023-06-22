package main

import (
	"fmt"
	"strings"

	"github.com/asdine/storm/v3/q"
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
		s.HasResource(),
	)
	if err != nil {
		return ctx.Error(err)
	}

	// récupérer les registres
	registers, err := Find[Register](
		s.HasResource(),
	)
	if err != nil {
		return ctx.Error(err)
	}

	// récupérer les posts
	topics, err := Find[Post](
		s.HasResource(),
		q.Eq("Parent", 0),
	)
	if err != nil {
		return ctx.Error(err)
	}

	fmt.Fprintf(tw, "LINKS\n")
	fmt.Fprintf(tw, "ID\tGroup\tAddress\tDescription\t\n")
	for _, l := range links {
		group := "public"
		if l.Group != "" {
			group = l.Group
		}
		fmt.Fprintf(tw, "%d\t%s\t%s\t%s\t\n", l.ID, group, l.Address, l.Description)
	}
	fmt.Fprintln(tw)

	fmt.Fprintf(tw, "REGISTERS\n")
	fmt.Fprintf(tw, "ID\tGroup\tDescription\tState\t\n")
	for _, r := range registers {
		group := "public"
		if r.Group != "" {
			group = r.Group
		}
		fmt.Fprintf(tw, "%d\t%s\t%s\t%s\t\n", r.ID, group, r.Description, r.State)
	}
	fmt.Fprintln(tw)

	fmt.Fprintf(tw, "FORUM TOPICS\n")
	fmt.Fprintf(tw, "ID\tGroup\tAuthor\tSubject\tState\t\n")
	for _, t := range topics {
		group := "public"
		if t.Group != "" {
			group = t.Group
		}
		fmt.Fprintf(tw, "%d\t%s\t%s\t%s\t\n", t.ID, group, t.Author, t.Subject)
	}

	tw.Flush()
	console.StartAlert()
	return ctx.Output(b.String())
}
