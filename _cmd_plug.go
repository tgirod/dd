package main

type PlugMsg struct{}

var plug = Cmd{
	Name:      "plug",
	ShortHelp: "active l'interface neuronale",
	Run:       Plug,
}

func Plug(ctx Context) any {
	console := ctx.Value("console").(*Console)
	res := ctx.Result()
	console.DNI = true
	res.Output = "interface neuronale directe activée"
	return res
}