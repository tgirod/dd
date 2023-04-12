package main

type QuitMsg struct{}

var quit = Cmd{
	Name:      "quit",
	ShortHelp: "ferme la connexion au serveur courant",
	Connected: true,
	Run: func(ctx Context, args []string) any {
		return QuitMsg{}
	},
}
