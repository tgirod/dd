package main

import (
	"fmt"
)

var evade = Cmd{
	name:       "evade",
	help:       "effectue une manoeuvre d'évasion pour gagner un peu de temps",
	connected:  true,
	identified: false,
	next: Select{
		name:   "mem",
		help:   "zone mémoire vers laquelle se déplacer",
		header: "liste des zones mémoire disponibles sur ce serveur",
		options: func(ctx Context) ([]Option, error) {
			console := ctx.Value("console").(*Console)
			opts := make([]Option, 0, len(console.Mem))
			for mem, available := range console.Mem {
				status := "OK"
				if !available {
					status = "INDISPONIBLE"
				}

				opts = append(opts, Option{
					value: mem,
					help:  status,
				})
			}
			return opts, nil
		},
		next: Run(Evade),
	},
}

func Evade(ctx Context) any {
	console := ctx.Value("console").(*Console)
	mem := ctx.Value("mem").(string)

	available, _ := console.Mem[mem]

	if !available {
		return ctx.Error(fmt.Errorf("%s : %w", mem, errMemUnavailable))
	}

	console.Mem[mem] = false

	console.StartAlert()

	return ctx.Result(nil, fmt.Sprintf("session relocalisée dans la zone mémoire %s", mem))
}
