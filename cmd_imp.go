package main

import (
	"fmt"
	"strings"
)

var imp = Cmd{
	name:       "imp",
	help:       "usurpe l'identité d'un utilisateur",
	connected:  true,
	identified: false,
	next: Select{
		name: "login",
		help: "compte utilisateur a usurper",
		options: func(ctx Context) ([]Option, error) {
			c := ctx.Console()
			users := c.Users()
			opts := make([]Option, len(users))
			for i, a := range users {
				opts[i] = Option{
					value: a.Login,
					help:  strings.Join(a.Groups(), " "),
				}
			}
			return opts, nil
		},
		next: Run(Imp),
	},
}

func Imp(ctx Context) any {
	c := ctx.Console()
	login := ctx.Value("login").(string)

	user, err := c.FindUser(login)
	if err != nil {
		ctx.Error(err)
	}
	c.Session.User = user

	c.StartAlert()

	return ctx.Output(fmt.Sprintf("usurpation du compte %s", login))
}
