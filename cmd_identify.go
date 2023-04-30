package main

import "fmt"

var identify = Cmd{
	Name:      "identify",
	ShortHelp: "validation d'identité avec le login/password",
	Args: []Arg{
		{
			Name:      "login",
			ShortHelp: "identifiant utilsateur",
			Type:      ShortArg,
		},
		{
			Name:      "password",
			ShortHelp: "mot de passe utilisateur",
			Type:      HiddenArg,
		},
	},
	Run: Identify,
}

func Identify(ctx Context) any {
	console := ctx.Value("console").(*Console)
	login := ctx.Value("login").(string)
	password := ctx.Value("password").(string)

	res := ctx.Result()
	if err := console.Identify(login, password); err != nil {
		res.Error = err
		return res
	}

	res.Output = fmt.Sprintf("vous êtes identifié en tant que %s", login)
	return res
}
