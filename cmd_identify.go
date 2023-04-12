package main

import "dd/ui/filler"

type IdentifyMsg struct {
	Login    string
	Password string
}

func (i IdentifyMsg) SetPassword(password string) filler.PasswordFiller {
	i.Password = password
	return i
}

func (i IdentifyMsg) GetPassword() string {
	return i.Password
}

var identify = Cmd{
	Name:      "identify",
	ShortHelp: "validation d'identité avec le login/password",
	Args: []Arg{
		{
			Name:      "login",
			ShortHelp: "identifiant utilsateur",
		},
	},
	Run: func(ctx Context, args []string) any {
		msg := IdentifyMsg{Login: args[0]}
		model := filler.New("entrez votre mot de passe", msg)
		return OpenModalMsg(model)
	},
}
