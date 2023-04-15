package main

var help = Cmd{
	Name:      "help",
	ShortHelp: "affiche l'aide pour une commande",
	Run:       Help,
}

func Help(ctx Context) any {
	res := ctx.Result()
	res.Output = ctx.Help(ctx.Args)
	return res
}
