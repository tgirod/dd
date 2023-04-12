package main

import (
	"dd/ui/filler"
	"fmt"
	"strconv"
	"strings"
)

type BalanceMsg struct{}

type PayMsg struct {
	To       string
	Amount   int
	Password string
}

func (p PayMsg) SetPassword(password string) filler.PasswordFiller {
	p.Password = password
	return p
}

func (p PayMsg) GetPassword() string {
	return p.Password
}

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
			Run: func(ctx Context, args []string) any {
				return BalanceMsg{}
			},
		},
		{
			Path:      []string{"yes"},
			Name:      "pay",
			ShortHelp: "effectue un transfert de monnaie",
			Args: []Arg{
				{
					Name:      "account",
					ShortHelp: "compte à créditer",
				},
				{
					Name:      "amount",
					ShortHelp: "montant à transférer",
				},
			},
			Run: func(ctx Context, args []string) any {
				to := args[0]
				amount, err := strconv.Atoi(args[1])
				if err != nil {
					return Result{
						Error: fmt.Errorf("montant : %w", errInvalidArgument),
					}
				}
				msg := PayMsg{To: to, Amount: amount}
				model := filler.New("entrez votre mot de passe", msg)
				return OpenModalMsg(model)
			},
		},
	},
}

func YesBalance(ctx Context, args []string) any {
	id := ctx.Console.Identity

	b := strings.Builder{}
	tw := tw(&b)
	fmt.Fprintf(tw, "Compte bancaire associé à l'identité %s\n", id.Login)
	fmt.Fprintf(tw, "Solde du compte :\t%d Y€S\t\n", id.Yes)
	tw.Flush()

	eval := ctx.Result()
	eval.Output = b.String()
	return eval
}
