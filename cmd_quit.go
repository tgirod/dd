package main

var quit = Cmd{
	name:      "quit",
	help:      "ferme la connexion au serveur courant",
	connected: true,
	next:      Run(Quit),
}

func Quit(ctx Context) any {
	console := ctx.Value("console").(*Console)

	console.Server = nil
	console.Identity = nil
	console.Account = nil
	console.Alert = false
	console.History.Clear()
	// BUG comment retirer les commandes de hack ?
	// console.Branch = baseCmds

	return ctx.Result(nil, "déconnexion effectuée")
}
