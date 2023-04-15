package main

var message = Cmd{
	Name:       "message",
	ShortHelp:  "consulter et envoyer des messages privés",
	Connected:  true,
	Identified: true,
	SubCmds: []Cmd{
		{
			Path:      []string{"message"},
			Name:      "new",
			ShortHelp: "lister les messages non lus",
		},
		{
			Path:      []string{"message"},
			Name:      "list",
			ShortHelp: "lister tous les messages",
		},
		{
			Path:      []string{"message"},
			Name:      "view",
			ShortHelp: "voir un message",
			Args: []Arg{
				{
					Name:      "id",
					ShortHelp: "index du message à consulter",
				},
			},
		},
		{
			Path:      []string{"message"},
			Name:      "send",
			ShortHelp: "écrire un message",
			Args: []Arg{
				{
					Name:      "recipient",
					ShortHelp: "destinataire du message",
				},
			},
		},
		{
			Path:      []string{"message"},
			Name:      "reply",
			ShortHelp: "répondre à un message",
			Args: []Arg{
				{
					Name:      "id",
					ShortHelp: "identifiant du message auquel répondre",
				},
			},
		},
	},
}
