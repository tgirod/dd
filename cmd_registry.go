package main

import (
	"fmt"
)

var registry = Cmd{
	name:      "registry",
	help:      "liste et manipule les registres du serveur",
	connected: true,
	next: Select{
		name: "id",
		help: "nom du registre",
		options: func(ctx Context) []Option {
			console := ctx.Value("console").(*Console)
			opts := make([]Option, len(console.Registers()))
			for i, r := range console.Registers() {
				opts[i] = Option{
					value: r.ID,
					help:  fmt.Sprintf("%s : %s", r.Description, r.State),
				}
			}
			return opts
		},
		next: Select{
			name: "state",
			help: "état à écrire dans le registre",
			options: func(ctx Context) []Option {
				console := ctx.Value("console").(*Console)
				id := ctx.Value("id").(int)
				reg, _ := console.Register(id) // FIXME ignore une erreur
				opts := make([]Option, len(reg.Options))
				for i, o := range reg.Options {
					opts[i] = Option{
						value: i,
						help:  o,
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
	state := ctx.Value("state").(int)

	reg, err := console.Register(id)
	if err != nil {
		return ctx.Error(err)
	}
	reg.State = reg.Options[state]

	return ctx.Output(fmt.Sprintf("nouvel état du registre : %s\n", reg.State))
}
