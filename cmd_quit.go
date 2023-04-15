package main

var quit = Cmd{
	Name:      "quit",
	ShortHelp: "ferme la connexion au serveur courant",
	Connected: true,
	Run:       Quit,
}

func Quit(ctx Context) any {
	res := ctx.Result()

	ctx.Server = nil
	ctx.Identity = nil
	ctx.Account = nil
	ctx.Alert = false
	ctx.History.Clear()
	// FIXME décharger les commandes de hack ?

	res.Output = "déconnexion effectuée"
	return res
}
