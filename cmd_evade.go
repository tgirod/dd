package main

import (
	"fmt"
	"strings"
)

var evade = Cmd{
	name:       "evade",
	help:       "effectue une manoeuvre d'évasion pour gagner un peu de temps",
	connected:  true,
	identified: false,
	next: Branch{
		name: "action",
		cmds: []Cmd{
			{
				name: "list",
				help: "liste les zones mémoires disponibles pour une évasion",
				next: Run(EvadeList),
			},
			{
				name: "move",
				help: "effectue la manoeuvre d'evasion vers une zone mémoire",
				next: Run(EvadeMove),
			},
		}},
}

func EvadeMove(ctx Context) any {
	console := ctx.Value("console").(*Console)
	zone := ctx.Value("zone").(string)

	available, ok := console.Mem[zone]
	if !ok {
		return ctx.Result(fmt.Errorf("%s : %w", zone, errMemNotFound), "")
	}

	if !available {
		return ctx.Result(fmt.Errorf("%s : %w", zone, errMemUnavailable), "")
	}

	console.Mem[zone] = false
	console.Countdown = console.Server.Scan
	return ctx.Result(nil, fmt.Sprintf("session relocalisée dans la zone mémoire %s", zone))
}

func EvadeList(ctx Context) any {
	console := ctx.Value("console").(*Console)

	b := strings.Builder{}
	tw := tw(&b)
	fmt.Fprintf(tw, "ZONE\tDISPONIBILITE\t\n")
	for addr, available := range console.Mem {
		if !available {
			fmt.Fprintf(tw, "%s\t%s\t\n", addr, "INDISPONIBLE")
		} else {
			fmt.Fprintf(tw, "%s\t%s\t\n", addr, "OK")
		}
	}
	tw.Flush()

	return ctx.Result(nil, b.String())
}