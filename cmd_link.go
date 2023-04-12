package main

import (
	"fmt"
	"strconv"
)

type LinkListMsg struct{}

type LinkMsg struct {
	Id int
}

var link = Cmd{
	Name:      "link",
	ShortHelp: "utilise les liens pour se connecter à un autre serveur",
	Connected: true,
	SubCmds: []Cmd{
		{
			Name:      "list",
			Path:      []string{"link"},
			ShortHelp: "affiche la liste des liens disponibles",
			Run: func(ctx Context, args []string) any {
				return LinkListMsg{}
			},
		},
		{
			Name:      "connect",
			Path:      []string{"link"},
			ShortHelp: "suit un lien vers un autre serveur",
			Args: []Arg{
				{
					Name:      "id",
					ShortHelp: "identifiant du lien à suivre",
				},
			},
			Run: func(ctx Context, args []string) any {
				// récupérer le lien
				id, err := strconv.Atoi(args[0])
				if err != nil {
					return Result{
						Error: fmt.Errorf("ID : %w", errInvalidArgument),
					}
				}
				return LinkMsg{id}
			},
		},
	},
}
