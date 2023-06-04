package main

import (
	"fmt"
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
			return ToOptions(users), nil
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
