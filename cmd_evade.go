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
			console := ctx.Console()
			return ToOptions(console.Mem), nil
		},
		next: Run(Evade),
	},
}

func Evade(ctx Context) any {
	console := ctx.Console()
	mem := ctx.Value("mem").(string)
	sess := console.Session

	for i, m := range sess.Mem {
		if m.Address == mem {
			if m.Used {
				return ctx.Error(fmt.Errorf("%s : %w", mem, errMemUnavailable))
			}
			// trouvé une zone mémoire
			sess.Mem[i].Used = false
			sess.Countdown = COUNTDOWN
			console.StartAlert()

			return ctx.Output(fmt.Sprintf("session relocalisée dans la zone mémoire %s", mem))
		}
	}

	return ctx.Error(fmt.Errorf("%s : %w", mem, errMemNotFound))
}
