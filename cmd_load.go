package main

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
		},
	},
	Run: func(ctx Context, args []string) any {
		code := args[0]
		return LoadMsg{code}
	},
}
