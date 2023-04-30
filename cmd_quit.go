package main

var quit = Cmd{
	Name:      "quit",
	ShortHelp: "ferme la connexion au serveur courant",
	Connected: true,
	Run:       Quit,
}

func Quit(ctx Context) any {
	console := ctx.Value("console").(*Console)
	res := ctx.Result()

	console.Server = nil
	console.Identity = nil
	console.Account = nil
	console.Alert = false
	console.History.Clear()
	// FIXME décharger les commandes de hack ?

	res.Output = "déconnexion effectuée"
	return res
}
