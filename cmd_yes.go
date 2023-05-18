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
					name: "identity",
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
			{
				name: "history",
				help: "historique des transactions",
				next: Run(YesHistory),
			},
		},
	},
}

func YesBalance(ctx Context) any {
	console := ctx.Value("console").(*Console)
	id := console.Identity

	bal, err := id.Balance()
	if err != nil {
		return ctx.Error(err)
	}

	b := strings.Builder{}
	fmt.Fprintf(&b, "Compte bancaire associé à l'identité %s\n", id.Login)
	fmt.Fprintf(&b, "Solde du compte : %d Y€S\n", bal)
	return ctx.Output(b.String())
}

func YesHistory(ctx Context) any {
	console := ctx.Console()
	transactions, err := console.Identity.Transactions()
	if err != nil {
		return ctx.Error(err)
	}

	b := strings.Builder{}
	tw := tw(&b)

	fmt.Fprintf(tw, "Débit\tCrédit\tOpération\tCommentaire\t\n")
	for _, t := range transactions {
		if t.From == console.Session.Identity.Login {
			// débit
			fmt.Fprintf(tw, "%d\t\t%s\t%s\t\n", t.Yes, t.To, t.Comment)
		} else {
			// crédit
			fmt.Fprintf(tw, "\t%d\t%s\t%s\t\n", t.Yes, t.From, t.Comment)
		}
	}
	tw.Flush()

	return ctx.Output(b.String())
}

func YesPay(ctx Context) any {
	console := ctx.Value("console").(*Console)
	to := ctx.Value("identity").(string)
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
