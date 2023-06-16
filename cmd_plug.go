package main

type PlugMsg struct{}

var plug = Cmd{
	name: "plug",
	help: "active l'interface neuronale",
	next: Run(Plug),
}

func Plug(ctx Context) any {
	console := ctx.Console()
	console.DNI = true
	return ctx.Result(nil, "interface neuronale directe activée")
}
