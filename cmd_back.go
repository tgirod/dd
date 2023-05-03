package main

import (
	"fmt"
	"strings"
)

var back = Cmd{
	name:       "back",
	help:       "quitte le serveur actuel et se reconnecte au serveur précédent",
	connected:  true,
	identified: false,
	next:       Run(Back),
}

func Back(ctx Context) any {
	console := ctx.Value("console").(*Console)

	if len(console.History) == 1 {
		return ctx.Result(errInvalidCommand, "")
	}

	// enlever le serveur actuel
	console.History.Pop()

	prev, _ := console.History.Peek()

	server, err := console.FindServer(prev.Address)
	if err != nil {
		return ctx.Result(err, "")
	}

	login := ""
	if console.Identity != nil {
		login = console.Identity.Login
	}

	account, err := server.CheckAccount(login)
	if err != nil {
		return ctx.Result(err, "")
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
	return ctx.Result(nil, b.String())
}
