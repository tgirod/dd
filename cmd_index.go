package main

type IndexMsg struct{}

var index = Cmd{
	Name:      "index",
	ShortHelp: "liste les services disponibles dans le serveur courant",
	Connected: true,
	Run: func(ctx Context, args []string) any {
		return IndexMsg{}
	},
}
