package main

import (
	"fmt"
	"strconv"
	"strings"
)

var message = Cmd{
	Name:       "message",
	ShortHelp:  "consulter et envoyer des messages privés",
	Connected:  true,
	Identified: true,
	SubCmds: []Cmd{
		{
			Path:      []string{"message"},
			Name:      "list",
			ShortHelp: "lister les messages",
			Run:       MessageList,
		},
		{
			Path:      []string{"message"},
			Name:      "read",
			ShortHelp: "lire un message",
			Run:       MessageRead,
			Args: []Arg{
				{
					Type:      MessageId,
					Name:      "index",
					ShortHelp: "index du message à lire",
				},
			},
		},
		{
			Path:      []string{"message"},
			Name:      "write",
			ShortHelp: "écrire un message",
			Run:       MessageWrite,
			Args: []Arg{
				{
					Name:      "recipient",
					ShortHelp: "destinataire du message",
					Type:      Text,
				},
				{
					Type:      Text,
					Name:      "subject",
					ShortHelp: "sujet du message",
				},
				{
					Type:      LongText,
					Name:      "content",
					ShortHelp: "contenu du message",
				},
			},
		},
		{
			Path:      []string{"message"},
			Name:      "reply",
			ShortHelp: "répondre à un message",
			Run:       MessageReply,
			Args: []Arg{
				{
					Name:      "id",
					ShortHelp: "identifiant du message auquel répondre",
					Type:      Text,
				},
			},
		},
	},
}

func MessageList(ctx Context) any {
	b := strings.Builder{}
	tw := tw(&b)

	fmt.Fprintf(tw, "liste de tous les messages :\n")
	for i, m := range ctx.Messages {
		fmt.Fprintf(tw, "%d\t%s\t\n", i, m.Subject)
	}
	tw.Flush()

	res := ctx.Result()
	res.Output = b.String()
	return res
}

func (m Message) Title() string       { return m.Sender }
func (m Message) Description() string { return m.Subject }
func (m Message) FilterValue() string { return m.Subject }

func MessageRead(ctx Context) any {
	res := ctx.Result()

	index, err := strconv.Atoi(ctx.Args[0])
	if err != nil {
		res.Error = errInvalidArgument
		return res
	}

	if index < 0 || index >= len(ctx.Messages) {
		res.Error = errInvalidArgument
		return res
	}

	b := strings.Builder{}

	msg := ctx.Messages[index]
	ctx.Messages[index].Opened = true

	fmt.Fprintf(&b, "De : %s\n", msg.Recipient)
	fmt.Fprintf(&b, "Sujet : %s\n", msg.Subject)
	fmt.Fprintln(&b, msg.Content)

	res.Output = b.String()
	return res
}

func MessageWrite(ctx Context) any {
	res := ctx.Result()

	recipient := ctx.Args[0]

	subject := ctx.Args[1]

	content := ctx.Args[2]

	// envoyer le message
	msg := Message{
		Recipient: recipient,
		Sender:    ctx.Identity.Login,
		Subject:   subject,
		Content:   content,
	}

	if err := ctx.MessageSend(msg); err != nil {
		res.Error = err
		return res
	}

	res.Output = fmt.Sprintf("message envoyé à %s", recipient)
	return res
}

func MessageReply(ctx Context) any {
	res := ctx.Result()

	index, err := strconv.Atoi(ctx.Args[0])
	if err != nil {
		res.Error = errInvalidArgument
		return res
	}

	if index < 0 || index >= len(ctx.Messages) {
		res.Error = errInvalidArgument
		return res
	}

	original := ctx.Messages[index]
	recipient := original.Sender
	subject := ctx.Args[1]
	content := ctx.Args[2]

	// envoyer le message
	msg := Message{
		Recipient: recipient,
		Sender:    ctx.Identity.Login,
		Subject:   subject,
		Content:   content,
	}

	if err := ctx.MessageSend(msg); err != nil {
		res.Error = err
		return res
	}

	res.Output = fmt.Sprintf("message envoyé à %s", recipient)
	return res
}
