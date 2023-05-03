package main

var help = Cmd{
	Name:      "help",
	ShortHelp: "affiche l'aide pour une commande",
	Run:       Help,
}

func Help(ctx Context) any {
	// FIXME
	//console := ctx.Value("console").(*Console)
	res := ctx.Result()
	//res.Output = console.Help(ctx.Args)
	return res
}
