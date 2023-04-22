package main

import (
	"fmt"
	"strconv"
	"strings"
)

var yes = Cmd{
	Name:       "yes",
	ShortHelp:  "effectuer des opérations bancaires",
	Connected:  true,
	Identified: true,
	SubCmds: []Cmd{
		{
			Path:      []string{"yes"},
			Name:      "balance",
			ShortHelp: "affiche le solde du compte",
			Run:       YesBalance,
		},
		{
			Path:      []string{"yes"},
			Name:      "pay",
			ShortHelp: "effectue un transfert de monnaie",
			Args: []Arg{
				{
					Name:      "account",
					ShortHelp: "compte à créditer",
					Type:      TextArg,
				},
				{
					Name:      "amount",
					ShortHelp: "montant à transférer",
					Type:      AmountArg,
				},
				{
					Name:      "password",
					ShortHelp: "mot de passe utilisateur",
					Type:      PasswordArg,
				},
			},
			Run: YesPay,
		},
	},
}

func YesBalance(ctx Context) any {
	id := ctx.Console.Identity

	b := strings.Builder{}
	tw := tw(&b)
	fmt.Fprintf(tw, "Compte bancaire associé à l'identité %s\n", id.Login)
	fmt.Fprintf(tw, "Solde du compte :\t%d Y€S\t\n", id.Yes)
	tw.Flush()

	res := ctx.Result()
	res.Output = b.String()
	return res
}

func YesPay(ctx Context) any {
	res := ctx.Result()

	to := ctx.Args[0]

	amount, err := strconv.Atoi(ctx.Args[1])
	if err != nil {
		res.Error = errInvalidArgument
		return res
	}

	password := ctx.Args[2]
	if _, err := ctx.Network.CheckIdentity(ctx.Identity.Login, password); err != nil {
		res.Error = err
		return res
	}

	from := ctx.Identity.Login
	if err := ctx.Network.Pay(from, to, amount); err != nil {
		res.Error = err
		return res
	}

	res.Output = fmt.Sprintf("transfert effectué")
	return res
}