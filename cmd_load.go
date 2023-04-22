package main

import "fmt"

type LoadMsg struct {
	Code string
}

var load = Cmd{
	Name:      "load",
	ShortHelp: "charge une nouvelle commande",
	Args: []Arg{
		{
			Name:      "code",
			ShortHelp: "code de la commande",
			Type:      TextArg,
		},
	},
	Run: Load,
}

func Load(ctx Context) any {
	res := ctx.Result()

	code := ctx.Args[0]
	command, ok := Hack[code]
	if !ok {
		res.Error = fmt.Errorf("%s : %w", code, errInvalidArgument)
		return res
	}

	ctx.Cmd.SubCmds = append(ctx.Cmd.SubCmds, command)
	res.Output = fmt.Sprintf("%s : commande chargée", command.Name)
	return res
}