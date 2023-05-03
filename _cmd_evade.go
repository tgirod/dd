package main

import (
	"fmt"
	"strings"
)

var evade = Cmd{
	Name:      "evade",
	ShortHelp: "effectue une manoeuvre d'évasion pour gagner un peu de temps",
	Connected: true,
	SubCmds: []Cmd{
		{
			Name:      "list",
			ShortHelp: "liste les zones mémoires disponibles pour une évasion",
			Run:       EvadeList,
		},
		{
			Name:      "move",
			ShortHelp: "effectue la manoeuvre d'evasion vers une zone mémoire",
			Args: []Arg{
				{
					Name:      "zone",
					ShortHelp: "zone mémoire pour l'évasion",
					Type:      ShortArg,
				},
			},
			Run: EvadeMove,
		},
	},
}

func EvadeMove(ctx Context) any {
	console := ctx.Value("console").(*Console)
	zone := ctx.Value("zone").(string)
	res := ctx.Result()

	available, ok := console.Mem[zone]
	if !ok {
		res.Error = fmt.Errorf("%s : %w", zone, errMemNotFound)
		return res
	}

	if !available {
		res.Error = fmt.Errorf("%s : %w", zone, errMemUnavailable)
		return res
	}

	console.Mem[zone] = false
	console.Countdown = console.Server.Scan
	res.Output = fmt.Sprintf("session relocalisée dans la zone mémoire %s", zone)
	return res
}

func EvadeList(ctx Context) any {
	console := ctx.Value("console").(*Console)
	res := ctx.Result()

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

	res.Output = b.String()
	return res
}