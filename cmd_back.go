package main

type BackMsg struct{}

var back = Cmd{
	Name:      "back",
	ShortHelp: "quitte le serveur actuel et se reconnecte au serveur précédent",
	Connected: true,
	Run: func(ctx Context, args []string) any {
		return BackMsg{}
	},
}
