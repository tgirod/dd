package main

import (
	"fmt"
)

var registry = Cmd{
	name:      "registry",
	help:      "liste et manipule les registres du serveur",
	connected: true,
	next: Select{
		name:   "id",
		help:   "nom du registre",
		header: "liste des registres disponibles dans ce serveur",
		options: func(ctx Context) []Option {
			console := ctx.Value("console").(*Console)
			opts := make([]Option, len(console.Registers))
			for i, r := range console.Registers {
				opts[i] = Option{
					value: i,
					help:  fmt.Sprintf("%s : %s", r.Description, r.State),
				}
			}
			return opts
		},
		next: Select{
			name:   "state",
			help:   "état à écrire dans le registre",
			header: "liste des états possibles pour ce registre",
			options: func(ctx Context) []Option {
				console := ctx.Value("console").(*Console)
				id := ctx.Value("id").(int)
				reg := console.Registers[id]
				opts := make([]Option, len(reg.Options))
				for i, o := range reg.Options {
					opts[i] = Option{
						value: o,
						help:  "",
					}
				}
				return opts
			},
			next: Run(RegistryEdit),
		},
	},
}

func RegistryEdit(ctx Context) any {
	console := ctx.Value("console").(*Console)
	id := ctx.Value("id").(int)
	state := ctx.Value("state").(string)

	reg := console.Registers[id]
	reg.State = state

	return ctx.Output(fmt.Sprintf("nouvel état du registre : %s\n", state))
}
