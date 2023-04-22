package main

import (
	"fmt"
	"strconv"
	"strings"
)

var link = Cmd{
	Name:      "link",
	ShortHelp: "utilise les liens pour se connecter à un autre serveur",
	Connected: true,
	SubCmds: []Cmd{
		{
			Name:      "list",
			Path:      []string{"link"},
			ShortHelp: "affiche la liste des liens disponibles",
			Run:       LinkConnect,
			Args: []Arg{
				{
					Name:      "id",
					ShortHelp: "identifiant du lien à suivre",
					Type:      LinkArg,
				},
			},
		},
	},
}

func (l Link) Title() string       { return l.Address }
func (l Link) Description() string { return l.Desc }
func (l Link) FilterValue() string { return l.Address }

func LinkConnect(ctx Context) any {
	res := ctx.Result()

	id, err := strconv.Atoi(ctx.Args[0])
	if err != nil {
		res.Error = errInvalidArgument
		return res
	}

	if id < 0 || id >= len(ctx.Server.Links) {
		res.Error = errInvalidArgument
		return res
	}

	address := ctx.Server.Links[id].Address
	if err := ctx.Console.Connect(address, false); err != nil {
		res.Error = err
		return res
	}

	ctx.History.Push(Link{address, ""})

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", ctx.Server.Address)
	fmt.Fprintf(&b, "%s\n", ctx.Server.Description)
	res.Output = b.String()
	return res
}
