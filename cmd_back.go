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
	result := ctx.Result()

	if len(ctx.History) == 1 {
		result.Error = errInvalidCommand
		return result
	}

	// enlever le serveur actuel
	ctx.History.Pop()

	prev, _ := ctx.History.Peek()

	server, err := ctx.FindServer(prev.Address)
	if err != nil {
		result.Error = err
		return result
	}

	login := ""
	if ctx.Identity != nil {
		login = ctx.Identity.Login
	}

	account, err := server.CheckAccount(login)
	if err != nil {
		result.Error = err
		return result
	}

	ctx.Console.Server = server
	ctx.Console.Account = account

	if ctx.Account != nil && ctx.Account.Backdoor {
		ctx.RemoveAccount(login)
		ctx.RemoveIdentity(login)
	}

	ctx.InitMem()

	b := strings.Builder{}
	fmt.Fprintf(&b, "connexion établie à l'adresse %s\n\n", ctx.Server.Address)
	result.Output = b.String()

	return result
}
