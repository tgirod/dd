package main

import (
	"fmt"
	"strings"
)

var back = Cmd{
	Name:      "back",
	ShortHelp: "quitte le serveur actuel et se reconnecte au serveur précédent",
	Connected: true,
	Run:       Back,
}

func Back(ctx Context) any {
	console := ctx.Value("console").(*Console)
	result := ctx.Result()

	if len(console.History) == 1 {
		result.Error = errInvalidCommand
		return result
	}

	// enlever le serveur actuel
	console.History.Pop()

	prev, _ := console.History.Peek()

	server, err := console.FindServer(prev.Address)
	if err != nil {
		result.Error = err
		return result
	}

	login := ""
	if console.Identity != nil {
		login = console.Identity.Login
	}

	account, err := server.CheckAccount(login)
	if err != nil {
		result.Error = err
		return result
	}

	console.Server = server
	console.Account = account

	if console.Account != nil && console.Account.Backdoor {
		console.RemoveAccount(login)
		console.RemoveIdentity(login)
	}

	console.InitMem()

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", console.Address)
	result.Output = b.String()

	return result
}
