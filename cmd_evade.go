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
			Path:      []string{"evade"},
			ShortHelp: "liste les zones mémoires disponibles pour une évasion",
			Run:       EvadeList,
		},
		{
			Name:      "move",
			Path:      []string{"evade"},
			ShortHelp: "effectue la manoeuvre d'evasion vers une zone mémoire",
			Args: []Arg{
				{
					Name:      "zone",
					ShortHelp: "zone mémoire pour l'évasion",
					Type:      Text,
				},
			},
			Run: EvadeMove,
		},
	},
}

func EvadeMove(ctx Context) any {
	res := ctx.Result()

	zone := ctx.Args[0]
	available, ok := ctx.Mem[zone]
	if !ok {
		res.Error = fmt.Errorf("%s : %w", zone, errMemNotFound)
		return res
	}

	if !available {
		res.Error = fmt.Errorf("%s : %w", zone, errMemUnavailable)
		return res
	}

	ctx.Mem[zone] = false
	ctx.Countdown = ctx.Server.Scan
	res.Output = fmt.Sprintf("session relocalisée dans la zone mémoire %s", zone)
	return res
}

func EvadeList(ctx Context) any {
	res := ctx.Result()

	b := strings.Builder{}
	tw := tw(&b)
	fmt.Fprintf(tw, "ZONE\tDISPONIBILITE\t\n")
	for addr, available := range ctx.Mem {
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