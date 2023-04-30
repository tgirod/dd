package main

import (
	"fmt"
	"strings"
)

var message = Cmd{
	Name:       "message",
	ShortHelp:  "consulter et envoyer des messages privés",
	Connected:  true,
	Identified: true,
	SubCmds: []Cmd{
		{
			Name:      "read",
			ShortHelp: "lire un message",
			Args: []Arg{
				{
					Type:      SelectNumberArg,
					Name:      "id",
					ShortHelp: "id du message à lire",
					Options: func(ctx Context) []Option {
						console := ctx.Value("console").(*Console)
						msgs := console.Identity.Messages
						opts := make([]Option, len(msgs))
						for i, m := range msgs {
							opts[i].Desc = fmt.Sprintf("%d -- %s -- %s", i, m.Sender, m.Subject)
							opts[i].Value = i
						}
						return opts
					},
				},
			},
			Run: MessageRead,
		},
		{
			Name:      "write",
			ShortHelp: "écrire un message",
			Args: []Arg{
				{
					Name:      "recipient",
					ShortHelp: "destinataire du message",
					Type:      ShortArg,
				},
				{
					Type:      ShortArg,
					Name:      "subject",
					ShortHelp: "sujet du message",
				},
				{
					Type:      LongArg,
					Name:      "content",
					ShortHelp: "contenu du message",
				},
			},
			Run: MessageWrite,
		},
		{
			Name:      "reply",
			ShortHelp: "répondre à un message",
			Args: []Arg{
				{
					Name:      "id",
					ShortHelp: "id du message auquel répondre",
					Type:      SelectNumberArg,
					Options: func(ctx Context) []Option {
						console := ctx.Value("console").(*Console)
						msgs := console.Identity.Messages
						opts := make([]Option, len(msgs))
						for i, m := range msgs {
							opts[i].Desc = fmt.Sprintf("%d -- %s -- %s", i, m.Sender, m.Subject)
							opts[i].Value = i
						}
						return opts
					},
				},
				{
					Type:      ShortArg,
					Name:      "subject",
					ShortHelp: "sujet du message",
				},
				{
					Type:      LongArg,
					Name:      "content",
					ShortHelp: "contenu du message",
				},
			},
			Run: MessageReply,
		},
	},
}

func MessageRead(ctx Context) any {
	console := ctx.Value("console").(*Console)
	id := ctx.Value("id").(int)
	res := ctx.Result()

	if id < 0 || id >= len(console.Identity.Messages) {
		res.Error = errInvalidArgument
		return res
	}
	msg := console.Identity.Messages[id]
	console.Messages[id].Opened = true

	b := strings.Builder{}

	fmt.Fprintf(&b, "ID : %d\n", id)
	fmt.Fprintf(&b, "De : %s\n", msg.Recipient)
	fmt.Fprintf(&b, "Sujet : %s\n", msg.Subject)
	fmt.Fprintln(&b, msg.Content)

	res.Output = b.String()
	return res
}

func MessageWrite(ctx Context) any {
	console := ctx.Value("console").(*Console)
	recipient := ctx.Value("recipient").(string)
	subject := ctx.Value("subject").(string)
	content := ctx.Value("content").(string)
	res := ctx.Result()

	// envoyer le message
	msg := Message{
		Recipient: recipient,
		Sender:    console.Identity.Login,
		Subject:   subject,
		Content:   content,
	}

	if err := console.MessageSend(msg); err != nil {
		res.Error = err
		return res
	}

	res.Output = fmt.Sprintf("message envoyé à %s", recipient)
	return res
}

func MessageReply(ctx Context) any {
	console := ctx.Value("console").(*Console)
	id := ctx.Value("id").(int)
	subject := ctx.Value("subject").(string)
	content := ctx.Value("content").(string)
	res := ctx.Result()

	if id < 0 || id >= len(console.Identity.Messages) {
		res.Error = errInvalidArgument
		return res
	}

	recipient := console.Identity.Messages[id].Sender

	// envoyer le message
	msg := Message{
		Recipient: recipient,
		Sender:    console.Identity.Login,
		Subject:   subject,
		Content:   content,
	}

	if err := console.MessageSend(msg); err != nil {
		res.Error = err
		return res
	}

	res.Output = fmt.Sprintf("message envoyé à %s", recipient)
	return res
}
