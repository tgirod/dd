package main

import (
	"fmt"
	"strings"
)

var link = Cmd{
	Name:      "link",
	ShortHelp: "utilise les liens pour se connecter à un autre serveur",
	Connected: true,
	Args: []Arg{
		{
			Type:      SelectNumberArg,
			Name:      "id",
			ShortHelp: "identifiant du lien",
			Options: func(ctx Context) []Option {
				console := ctx.Value("console").(*Console)
				links := console.Server.Links
				opts := make([]Option, len(links))
				for i, l := range links {
					opts[i].Desc = fmt.Sprintf("%d -- %s", i, l.Desc)
					opts[i].Value = i
				}
				return opts
			},
		},
	},
	Run: LinkCmd,
}

func LinkCmd(ctx Context) any {
	console := ctx.Value("console").(*Console)
	res := ctx.Result()

	id := ctx.Value("id").(int)

	if id < 0 || id >= len(console.Server.Links) {
		res.Error = errInvalidArgument
		return res
	}

	link := console.Server.Links[id]

	if err := console.Connect(link.Address, false); err != nil {
		res.Error = err
		return res
	}

	console.History.Push(link)

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", console.Server.Address)
	fmt.Fprintf(&b, "%s\n", console.Server.Description)
	res.Output = b.String()
	return res
}
