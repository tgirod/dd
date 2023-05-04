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
					name: "id",
					help: "id du message à lire",
					options: func(ctx Context) []Option {
						console := ctx.Value("console").(*Console)
						msgs := console.Identity.Messages
						opts := make([]Option, len(msgs))
						for i, m := range msgs {
							opts[i].value = i
							opts[i].help = fmt.Sprintf("%d -- %s -- %s", i, m.Sender, m.Subject)
						}
						return opts
					},
					next: Run(MessageRead),
				},
			},
			{
				name: "write",
				help: "écrire un message",
				next: String{
					name: "recipient",
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
					name: "id",
					help: "id du message auquel répondre",
					options: func(ctx Context) []Option {
						console := ctx.Value("console").(*Console)
						msgs := console.Identity.Messages
						opts := make([]Option, len(msgs))
						for i, m := range msgs {
							opts[i].value = i
							opts[i].help = fmt.Sprintf("%d -- %s -- %s", i, m.Sender, m.Subject)
						}
						return opts
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
	console := ctx.Value("console").(*Console)
	id := ctx.Value("id").(int)

	msg := console.Identity.Messages[id]
	console.Messages[id].Opened = true

	b := strings.Builder{}

	fmt.Fprintf(&b, "ID : %d\n", id)
	fmt.Fprintf(&b, "De : %s\n", msg.Recipient)
	fmt.Fprintf(&b, "Sujet : %s\n", msg.Subject)
	fmt.Fprintln(&b, msg.Content)

	return ctx.Result(nil, b.String())
}

func MessageWrite(ctx Context) any {
	console := ctx.Value("console").(*Console)
	recipient := ctx.Value("recipient").(string)
	subject := ctx.Value("subject").(string)
	content := ctx.Value("content").(string)

	// envoyer le message
	msg := Message{
		Recipient: recipient,
		Sender:    console.Identity.Login,
		Subject:   subject,
		Content:   content,
	}

	if err := console.MessageSend(msg); err != nil {
		return ctx.Result(err, "")
	}

	return ctx.Result(nil, fmt.Sprintf("message envoyé à %s", recipient))
}

func MessageReply(ctx Context) any {
	console := ctx.Value("console").(*Console)
	id := ctx.Value("id").(int)
	subject := ctx.Value("subject").(string)
	content := ctx.Value("content").(string)

	recipient := console.Identity.Messages[id].Sender

	// envoyer le message
	msg := Message{
		Recipient: recipient,
		Sender:    console.Identity.Login,
		Subject:   subject,
		Content:   content,
	}

	if err := console.MessageSend(msg); err != nil {
		return ctx.Result(err, "")
	}

	return ctx.Result(nil, fmt.Sprintf("message envoyé à %s", recipient))
}
