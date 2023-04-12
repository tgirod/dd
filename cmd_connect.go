package main

type ConnectMsg struct {
	Address string
}

var connect = Cmd{
	Name:      "connect",
	ShortHelp: "établit la connexion avec un serveur",
	Args: []Arg{
		{
			Name:      "address",
			ShortHelp: "adresse sur serveur auquel se connecter",
		},
	},
	Run: func(ctx Context, args []string) any {
		return ConnectMsg{args[0]}
	},
}
