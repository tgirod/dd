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
			Type:      HiddenArg,
		},
	},
	Run: Load,
}

func Load(ctx Context) any {
	console := ctx.Value("console").(*Console)
	code := ctx.Value("code").(string)
	res := ctx.Result()

	command, ok := Hack[code]
	if !ok {
		res.Error = fmt.Errorf("%s : %w", code, errInvalidArgument)
		return res
	}

	console.Cmd.SubCmds = append(console.Cmd.SubCmds, command)
	res.Output = fmt.Sprintf("%s : commande chargée", command.Name)
	return res
}