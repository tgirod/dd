package main

import (
	"fmt"
	"strings"
)

var yes = Cmd{
	name:       "yes",
	help:       "effectuer des opérations bancaires",
	connected:  true,
	identified: true,
	next: Branch{
		name: "action",
		cmds: []Cmd{
			{
				name: "balance",
				help: "affiche le solde du compte",
				next: Run(YesBalance),
			},
			{
				name: "pay",
				help: "effectue un transfert de monnaie",
				next: String{
					name: "account",
					help: "compte à créditer",
					next: Number{
						name: "amount",
						help: "montant à transférer",
						next: Hidden{
							name: "password",
							help: "mot de passe du compte",
							next: Run(YesPay),
						},
					},
				},
			},
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

	return ctx.Result(nil, b.String())
}

func YesPay(ctx Context) any {
	console := ctx.Value("console").(*Console)
	to := ctx.Value("account").(string)
	amount := ctx.Value("amount").(int)
	password := ctx.Value("password").(string)

	if _, err := CheckIdentity(console.Identity.Login, password); err != nil {
		return ctx.Result(err, "")
	}

	from := console.Identity.Login
	if err := Pay(from, to, amount); err != nil {
		return ctx.Result(err, "")
	}

	return ctx.Result(nil, fmt.Sprintf("transfert effectué"))
}