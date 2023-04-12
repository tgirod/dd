package main

type RegistrySearchMsg struct {
	Name string
}

type RegistryEditMsg struct {
	Name string
}

var registry = Cmd{
	Name:      "registry",
	ShortHelp: "liste et manipule les registres du serveur",
	Connected: true,
	SubCmds: []Cmd{
		{
			Name:      "search",
			Path:      []string{"registry"},
			ShortHelp: "recherche dans les registres",
			Args: []Arg{
				{
					Name:      "prefix",
					ShortHelp: "préfixe du nom du registre",
				},
			},
			Run: func(ctx Context, args []string) any {
				return RegistrySearchMsg{args[0]}
			},
		},
		{
			Name:      "edit",
			Path:      []string{"registry"},
			ShortHelp: "modifie un registre",
			Args: []Arg{
				{
					Name:      "name",
					ShortHelp: "nom du registre à modifier",
				},
			},
			Run: func(ctx Context, args []string) any {
				return RegistryEditMsg{args[0]}
			},
		},
	},
}
