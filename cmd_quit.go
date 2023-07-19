package main

var quit = Cmd{
	name:      "quit",
	help:      "ferme la connexion et efface les programmes chargés",
	connected: false,
	next:      Run(Quit),
}

func Quit(ctx Context) any {
	console := ctx.Console()
	console.DNI = false
	console.Disconnect()

	return ctx.Result(nil, "déconnexion effectuée")
}
