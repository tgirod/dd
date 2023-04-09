package main

import (
	"dd/ui/filler"
	"fmt"
	"strconv"
)

type MessageNewMsg struct {
	Recipient string
	Subject   string
	Content   string
}

type MessageListMsg struct{}

type MessageViewMsg struct {
	Index int
}

type MessageSendMsg struct {
	Recipient string // destinataire du message
	Subject   string
	Content   string
}

func (m MessageSendMsg) SetSubject(subject string) filler.SubjectFiller {
	m.Subject = subject
	return m
}

func (m MessageSendMsg) GetSubject() string {
	return m.Subject
}

func (m MessageSendMsg) SetContent(content string) filler.ContentFiller {
	m.Content = content
	return m
}

func (m MessageSendMsg) GetContent() string {
	return m.Content
}

type MessageReplyMsg struct {
	Index int // identifiant du message auquel on répond
}

var message = Cmd{
	Name:       "message",
	ShortHelp:  "consulter et envoyer des messages privés",
	Connected:  true,
	Identified: true,
	SubCmds: []Cmd{
		{
			Path:      []string{"message"},
			Name:      "new",
			ShortHelp: "lister les messages non lus",
			Parse: func(args []string) any {
				return MessageNewMsg{}
			},
		},
		{
			Path:      []string{"message"},
			Name:      "list",
			ShortHelp: "lister tous les messages",
			Parse: func(args []string) any {
				return MessageListMsg{}
			},
		},
		{
			Path:      []string{"message"},
			Name:      "view",
			ShortHelp: "voir un message",
			Args: []Arg{
				{
					Name:      "id",
					ShortHelp: "index du message à consulter",
				},
			},
			Parse: func(args []string) any {
				id, err := strconv.Atoi(args[0])
				if err != nil {
					return Eval{
						Error: fmt.Errorf("ID : %w", errInvalidArgument),
					}
				}
				return MessageViewMsg{
					Index: id,
				}
			},
		},
		{
			Path:      []string{"message"},
			Name:      "send",
			ShortHelp: "écrire un message",
			Args: []Arg{
				{
					Name:      "recipient",
					ShortHelp: "destinataire du message",
				},
			},
			Parse: func(args []string) any {
				msg := MessageSendMsg{
					Recipient: args[0],
				}
				model := filler.New("saisissez votre message", msg)
				return OpenModalMsg(model)
			},
		},
		{
			Path:      []string{"message"},
			Name:      "reply",
			ShortHelp: "répondre à un message",
			Args: []Arg{
				{
					Name:      "id",
					ShortHelp: "identifiant du message auquel répondre",
				},
			},
			Parse: func(args []string) any {
				id, err := strconv.Atoi(args[0])
				if err != nil {
					return Eval{
						Error: fmt.Errorf("ID : %w", errInvalidArgument),
					}
				}

				return MessageReplyMsg{
					Index: id,
				}
			},
		},
	},
}
