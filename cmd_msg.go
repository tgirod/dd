package main

import (
	"fmt"
	"strings"
)

var message = Cmd{
	name:       "message",
	help:       "consulter et envoyer des messages privés",
	connected:  true,
	identified: true,
	next: Branch{
		name: "action",
		cmds: []Cmd{
			{
				name: "read",
				help: "lire un message",
				next: Select{
					name:   "id",
					help:   "id du message à lire",
					header: "liste des messages reçus et envoyés",
					options: func(ctx Context) ([]Option, error) {
						console := ctx.Console()
						msgs := console.Identity.Messages()
						return ToOptions(msgs), nil
					},
					next: Run(MessageRead),
				},
			},
			{
				name: "write",
				help: "écrire un message",
				next: String{
					name: "to",
					help: "destinataire du message",
					next: Text{
						name: "subject",
						help: "sujet du message",
						next: LongText{
							name: "content",
							help: "contenu du message",
							next: Run(MessageWrite),
						},
					},
				},
			},
			{
				name: "reply",
				help: "répondre à un message",
				next: Select{
					name:   "id",
					help:   "id du message auquel répondre",
					header: "liste des messages reçus et envoyés",
					options: func(ctx Context) ([]Option, error) {
						console := ctx.Console()
						msgs := console.Identity.Messages()
						return ToOptions(msgs), nil
					},
					next: LongText{
						name: "content",
						help: "contenu du message",
						next: Run(MessageReply),
					},
				},
			},
		},
	},
}

func MessageRead(ctx Context) any {
	console := ctx.Console()
	id := ctx.Value("id").(int)

	msg, err := console.Identity.Message(id)
	if err != nil {
		return ctx.Error(err)
	}

	// marquer le message comme lu
	if !msg.Opened {
		msg.Opened = true
		msg, err = Save(msg)
		if err != nil {
			return ctx.Error(err)
		}
	}

	return ctx.Result(nil, msg.String())
}

func (m Message) String() string {
	b := strings.Builder{}

	fmt.Fprintf(&b, "ID : %d\n", m.ID)
	fmt.Fprintf(&b, "De : %s\n", m.From)
	fmt.Fprintf(&b, "À : %s\n", m.To)
	fmt.Fprintf(&b, "Date : %s\n", m.Date)
	fmt.Fprintf(&b, "Sujet : %s\n", m.Subject)
	fmt.Fprintln(&b, m.Content)

	return b.String()
}

func MessageWrite(ctx Context) any {
	console := ctx.Console()
	identity := console.Identity
	to := ctx.Value("to").(string)
	subject := ctx.Value("subject").(string)
	content := ctx.Value("content").(string)

	if _, err := identity.Send(to, subject, content); err != nil {
		return ctx.Result(err, "")
	}

	return ctx.Result(nil, fmt.Sprintf("message envoyé à %s", to))
}

func MessageReply(ctx Context) any {
	console := ctx.Console()
	identity := console.Identity
	id := ctx.Value("id").(int)
	original, err := console.Message(id)
	if err != nil {
		return ctx.Error(err)
	}

	subject := fmt.Sprintf("Re: %s", original.Subject)
	content := ctx.Value("content").(string)

	if _, err := identity.Send(original.From, subject, content); err != nil {
		return ctx.Result(err, "")
	}

	return ctx.Result(nil, fmt.Sprintf("message envoyé à %s", original.From))
}
