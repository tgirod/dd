package main

type HelpMsg struct {
	Args []string
}

var help = Cmd{
	Name:      "help",
	ShortHelp: "affiche l'aide",
	Run: func(ctx Context, args []string) any {
		return HelpMsg{args}
	},
}
