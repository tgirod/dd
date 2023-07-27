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
		help:   "identifiant du registre",
		header: "liste des registres disponibles dans ce serveur",
		options: func(ctx Context) ([]Option, error) {
			console := ctx.Console()
			return ToOptions(console.Registers(console.User)), nil
		},
		next: Branch{
			name: "",
			cmds: []Cmd{
				{
					name: "show",
					help: "affiche l'état du registre",
					next: Run(RegistryShow),
				},
				{
					name: "edit",
					help: "édite l'état du registre",
					next: Select{
						name:   "state",
						help:   "état à écrire dans le registre",
						header: "liste des états possibles pour ce registre",
						options: func(ctx Context) ([]Option, error) {
							console := ctx.Console()
							id := ctx.Value("id").(int)
							reg, err := console.Register(id, console.User)
							if err != nil {
								return []Option{}, err
							}
							return ToOptions(reg.Options), nil
						},
						next: Run(RegistryEdit),
					},
				},
			},
		},
	},
}

func RegistryEdit(ctx Context) any {
	console := ctx.Console()
	id := ctx.Value("id").(int)
	state := ctx.Value("state").(RegisterState)

	reg, err := console.Register(id, console.User)
	if err != nil {
		return ctx.Error(err)
	}
	reg.State = state
	reg, err = Save(reg)
	if err != nil {
		return ctx.Error(err)
	}

	return ctx.Output(fmt.Sprintf("nouvel état du registre : %s\n", reg.State))
}

func RegistryShow(ctx Context) any {
	console := ctx.Console()
	id := ctx.Value("id").(int)

	reg, err := console.Register(id, console.User)
	if err != nil {
		return ctx.Error(err)
	}

	return ctx.Output(fmt.Sprintf("état du registre %d : %s\n", reg.ID, reg.State))
}
