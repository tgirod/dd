package main

type PlugMsg struct{}

var plug = Cmd{
	Name:      "plug",
	ShortHelp: "active l'interface neuronale",
	Run:       Plug,
}

func Plug(ctx Context) any {
	res := ctx.Result()
	ctx.DNI = true
	res.Output = "interface neuronale directe activée"
	return res
}