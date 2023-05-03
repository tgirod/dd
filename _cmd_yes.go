package main

import (
	"fmt"
	"strings"
)

var yes = Cmd{
	Name:       "yes",
	ShortHelp:  "effectuer des opérations bancaires",
	Connected:  true,
	Identified: true,
	SubCmds: []Cmd{
		{
			Name:      "balance",
			ShortHelp: "affiche le solde du compte",
			Run:       YesBalance,
		},
		{
			Name:      "pay",
			ShortHelp: "effectue un transfert de monnaie",
			Args: []Arg{
				{
					Name:      "account",
					ShortHelp: "compte à créditer",
					Type:      ShortArg,
				},
				{
					Name:      "amount",
					ShortHelp: "montant à transférer",
					Type:      NumberArg,
				},
				{
					Name:      "password",
					ShortHelp: "mot de passe utilisateur",
					Type:      HiddenArg,
				},
			},
			Run: YesPay,
		},
	},
}

func YesBalance(ctx Context) any {
	console := ctx.Value("console").(*Console)
	id := console.Identity

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
	console := ctx.Value("console").(*Console)
	to := ctx.Value("account").(string)
	amount := ctx.Value("amount").(int)
	password := ctx.Value("password").(string)
	res := ctx.Result()

	if _, err := console.CheckIdentity(console.Identity.Login, password); err != nil {
		res.Error = err
		return res
	}

	from := console.Identity.Login
	if err := console.Pay(from, to, amount); err != nil {
		res.Error = err
		return res
	}

	res.Output = fmt.Sprintf("transfert effectué")
	return res
}