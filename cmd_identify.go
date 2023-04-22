package main

import "fmt"

var identify = Cmd{
	Name:      "identify",
	ShortHelp: "validation d'identité avec le login/password",
	Args: []Arg{
		{
			Name:      "login",
			ShortHelp: "identifiant utilsateur",
			Type:      LoginArg,
		},
		{
			Name:      "password",
			ShortHelp: "mot de passe utilisateur",
			Type:      PasswordArg,
		},
	},
	Run: Identify,
}

func Identify(ctx Context) any {
	login := ctx.Args[0]
	password := ctx.Args[1]

	res := ctx.Result()
	if err := ctx.Console.Identify(login, password); err != nil {
		res.Error = err
		return res
	}

	res.Output = fmt.Sprintf("vous êtes identifié en tant que %s", login)
	return res
}
