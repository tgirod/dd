package main

import "fmt"

var identify = Cmd{
	name:       "identify",
	help:       "validation d'identité avec le login/password",
	connected:  false,
	identified: false,
	next: String{
		name: "login",
		help: "identifiant utilisateur",
		next: Hidden{
			name: "password",
			help: "mot de passe utilisateur",
			next: Run(Identify),
		},
	},
}

func Identify(ctx Context) any {
	console := ctx.Console()
	login := ctx.Value("login").(string)
	password := ctx.Value("password").(string)

	if err := console.Identify(login, password); err != nil {
		return ctx.Result(err, "")
	}

	return ctx.Result(nil, fmt.Sprintf("vous êtes identifié en tant que %s", login))
}
