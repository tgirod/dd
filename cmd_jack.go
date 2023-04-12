package main

import (
	"dd/ui/loader"
	"fmt"
	"strconv"
	"time"
)

type JackMsg struct {
	Id int
}

var jack = Cmd{
	Name:      "jack",
	ShortHelp: "force l'accès à un lien",
	Connected: true,
	Args: []Arg{
		{
			Name:      "id",
			ShortHelp: "identifiant du lien",
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
		msg := JackMsg{id}
		model := loader.New(
			msg,
			3*time.Second,
			[]string{
				"recherche d'une faille",
				"exploit",
				"accès",
			},
		)
		return OpenModalMsg(model)
	},
}
