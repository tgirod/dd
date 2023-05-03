package main

import "fmt"

type LoadMsg struct {
	Code string
}

var load = Cmd{
	name: "load",
	help: "charge une nouvelle commande",
	next: String{
		name: "code",
		help: "code de la commande",
		next: Run(Load),
	},
}

func Load(ctx Context) any {
	console := ctx.Value("console").(*Console)
	code := ctx.Value("code").(string)

	command, ok := Hack[code]
	if !ok {
		return ctx.Result(fmt.Errorf("%s : %w", code, errInvalidArgument), "")
	}

	console.Branch.cmds = append(console.Branch.cmds, command)
	return ctx.Result(nil, fmt.Sprintf("%s : commande chargée", command.name))
}