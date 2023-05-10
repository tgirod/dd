package main

var quit = Cmd{
	name:      "quit",
	help:      "ferme la connexion au serveur courant",
	connected: true,
	next:      Run(Quit),
}

func Quit(ctx Context) any {
	console := ctx.Value("console").(*Console)
	console.DNI = false
	console.Disconnect()

	return ctx.Result(nil, "déconnexion effectuée")
}
