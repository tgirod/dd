package main

type DoorMsg struct{}

var door = Cmd{
	Name:      "door",
	ShortHelp: "créé une backdoor dans le serveur",
	Connected: true,
	Run: func(ctx Context, args []string) any {
		return DoorMsg{}
	},
}
