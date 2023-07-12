package main

import (
	"fmt"
	"strings"

	"github.com/asdine/storm/v3/q"
)

// Commands for Monitoring. Should only be available in Monitor

// _id manipuler identités *****************************************************
var sudo_id = Cmd{
	name: "_id",
	help: "manipuler les identités (list, take)",
	next: Branch{
		name: "_id_action",
		cmds: []Cmd{
			{
				name: "list",
				help: "lister les identités",
				next: Run(SudoListIdentity),
			},
			{
				name: "take",
				help: "choisi une identité",
				next: String{
					name: "login",
					help: "login de l'identité",
					next: Run(SudoTakeIdentity),
				},
			},
		},
	},
}
func SudoListIdentity(ctx Context) any {
	identities, err := Identities()
	if err != nil {
		return ctx.Error(err)
	}

	b := strings.Builder{}
	for _, i := range identities {
		fmt.Fprintf(&b, "%+v\n", i)
	}
	return ctx.Output(b.String())
}
func SudoTakeIdentity(ctx Context) any {
	wantedLogin := ctx.Value("login").(string)

	// Look for the corresponding password
	identity, err := FindIdentity(wantedLogin)
	if err != nil {
		return ctx.Error(err )
	}
	if err = ctx.Console().Identify(wantedLogin, identity.Password); err != nil {
		return ctx.Error(err)
	}

	return ctx.Result(nil,
		fmt.Sprintf("vous êtes identifié en tant que %s", wantedLogin))
}

// _msg manipuler les messages *************************************************
var sudo_msg = Cmd{
	name: "_msg",
	help: "manipuler les messages (from, to)",
	next: Branch{
		name: "",
		cmds: []Cmd{
			{
				name: "from",
				help: "liste messages envoyés par un utilisateur",
				next: Select{
					name:   "from",
					help:   "auteur des messages",
					header: "liste des auteurs",
					options: func(ctx Context) ([]Option, error) {
						identities, err := Identities()
						if err != nil {
							return []Option{}, err
						}
						return ToOptions(identities), nil
					},
					next: Run(SudoMsgFrom),
				},
			},
			{
				name: "to",
				help: "liste messages envoyés à un utilisateur",
				next: Select{
					name:   "to",
					help:   "destinataire des messages",
					header: "liste des destinataires",
					options: func(ctx Context) ([]Option, error) {
						identities, err := Identities()
						if err != nil {
							return []Option{}, err
						}
						return ToOptions(identities), nil
					},
					next: Run(SudoMsgTo),
				},
			},
		},
	},
}
func SudoMsgFrom(ctx Context) any {
	from := ctx.Value("from").(string)
	messages, err := Find[Message](
		q.Eq("From", from),
	)
	if err != nil {
		return ctx.Error(err)
	}

	b := strings.Builder{}
	fmt.Fprintf(&b, "messages envoyés par %s : %d\n", from, len(messages))

	for _, m := range messages {
		fmt.Fprintln(&b, m.String())
	}

	return ctx.Output(b.String())
}

func SudoMsgTo(ctx Context) any {
	to := ctx.Value("to").(string)
	messages, err := Find[Message](
		q.Eq("To", to),
	)
	if err != nil {
		return ctx.Error(err)
	}

	b := strings.Builder{}
	fmt.Fprintf(&b, "messages envoyés à %s : %d\n", to, len(messages))

	for _, m := range messages {
		fmt.Fprintln(&b, m.String())
	}

	return ctx.Output(b.String())
}

// _yes manipuler les yes :o) **************************************************
var sudo_yes = Cmd{
	name: "_yes",
	help: "manipuler les YES (balande, history)",
	next: Branch{
		name: "action",
		cmds: []Cmd{
			{
				name: "list",
				help: "afficher le solde de tous les comptes",
				next: Run(SudoYesList),
			},
			{
				name: "history",
				help: "historique des transaction d'un user",
				next: Select{
					name:   "login",
					help:   "proprio du compte",
					header: "liste des propriétaires de comptes",
					options: func(ctx Context) ([]Option, error) {
						identities, err := Identities()
						if err != nil {
							return []Option{}, err
						}
						return ToOptions(identities), nil
					},
					next: Run(SudoYesHistory),
				},
			},
		},
	},
}
func SudoYesList(ctx Context) any {
	identities, err := Identities()
	if err != nil {
		return ctx.Error(err)
	}

	s := strings.Builder{}
	fmt.Fprintf(&s, underline.Render("ID                           Solde"))
	fmt.Fprintf(&s, "\n")
	tw := tw(&s)
	fmt.Fprintf(tw, "\n")

	for _, id := range identities {
		bal, err := id.Balance()
		if err != nil {
			return ctx.Error(err)
		}
		fmt.Fprintf(tw, "%s\t\t\t\t%d\t\n", id.Login, bal)
	}
	tw.Flush()
	return ctx.Output(s.String())
}
func SudoYesHistory(ctx Context) any {
	login := ctx.Value("login").(string)
	id, err := FindIdentity(login)
	if err != nil {
		return ctx.Error(err)
	}

	transactions, err := id.Transactions()
	if err != nil {
		return ctx.Error(err)
	}

	b := strings.Builder{}
	tw := tw(&b)

	fmt.Fprintf(tw, "Débit\tCrédit\tOpération\tCommentaire\t\n")
	for _, t := range transactions {
		if t.From == id.Login {
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
