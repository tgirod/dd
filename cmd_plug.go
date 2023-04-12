package main

type PlugMsg struct{}

var plug = Cmd{
	Name:      "plug",
	ShortHelp: "active l'interface neuronale",
	Run: func(ctx Context, args []string) any {
		return PlugMsg{}
	},
}
