package main

type EvadeListMsg struct{}

type EvadeMsg struct {
	Zone string
}

var evade = Cmd{
	Name:      "evade",
	ShortHelp: "effectue une manoeuvre d'évasion pour gagner un peu de temps",
	Connected: true,
	SubCmds: []Cmd{
		{
			Name:      "list",
			Path:      []string{"evade"},
			ShortHelp: "liste les zones mémoires disponibles pour une évasion",
			Run: func(ctx Context, args []string) any {
				return EvadeListMsg{}
			},
		},
		{
			Name:      "move",
			Path:      []string{"evade"},
			ShortHelp: "effectue la manoeuvre d'evasion vers une zone mémoire",
			Args: []Arg{
				{
					Name:      "zone",
					ShortHelp: "zone mémoire pour l'évasion",
				},
			},
			Run: func(ctx Context, args []string) any {
				return EvadeMsg{args[0]}
			},
		},
	},
}
