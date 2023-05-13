package main

import (
	"fmt"
	"strings"
	"time"
)

func topicList(ctx Context) ([]Option, error) {
	console := ctx.Value("console").(*Console)
	topics := console.Server.Topics(console.User)
	opts := make([]Option, 0, len(topics))
	for _, t := range topics {
		opts = append(opts, Option{
			help:  fmt.Sprintf("%s\t%s\t%s\t", t.Date.Format(time.DateTime), t.Author, t.Subject),
			value: t.ID,
		})
	}
	return opts, nil
}

func postList(ctx Context) ([]Option, error) {
	console := ctx.Value("console").(*Console)
	topic := ctx.Value("topic").(int)
	posts := console.Server.RecReplies(topic, console.User)
	opts := make([]Option, 0, len(posts))
	for _, p := range posts {
		if p.Parent == topic {
			opts = append(opts, Option{
				help:  fmt.Sprintf("%s\t%s\t%s\t", p.Date.Format(time.DateTime), p.Author, p.Subject),
				value: p.ID,
			})
		}
	}
	return opts, nil
}

func threadList(ctx Context) ([]Option, error) {
	console := ctx.Console()
	topic := ctx.Value("topic").(int)
	// récupérer le post racine
	root, err := console.Server.Post(topic, console.User)
	if err != nil {
		return []Option{}, err
	}
	// récupérer le thread
	thread, err := console.Server.Thread(root, console.User)
	return thread.ToOptions(""), nil
}

var rep = strings.NewReplacer("├", "│", "└", " ", "─", " ")

func (t Thread) ToOptions(prefix string) []Option {
	// afficher le message à la racine du thread
	opts := []Option{
		{
			value: t.Post.ID,
			help:  fmt.Sprintf("%s\t%s\t%s%s\t", t.Date.Format(time.DateTime), t.Author, prefix, t.Subject),
		},
	}

	// pour les messages qui suivent la racine, le préfixe change pour poursuivre les traits
	prefix = rep.Replace(prefix)

	// appel récursif sur chaque réponse
	for i, reply := range t.Replies {
		if i < len(t.Replies)-1 {
			opts = append(opts, reply.ToOptions(prefix+"├─ ")...)
		} else {
			opts = append(opts, reply.ToOptions(prefix+"└─ ")...)
		}
	}
	return opts
}

var forum = Cmd{
	name:       "forum",
	help:       "participer au forum du serveur",
	connected:  true,
	identified: true,
	next: Branch{
		name: "action",
		cmds: []Cmd{
			{
				name: "read",
				help: "lire les posts",
				next: Select{
					name:    "topic",
					help:    "sujet de discussion",
					header:  "liste des sujets de discussions sur ce serveur",
					options: topicList,
					next: Select{
						name:    "post",
						help:    "message dans la discussion",
						header:  "liste des messages dans ce sujet de discussion",
						options: threadList,
						next:    Run(PostRead),
					},
				},
			},
			{
				name: "write",
				help: "ouvrir un nouveau sujet",
				next: Text{
					name: "subject",
					help: "sujet du post",
					next: LongText{
						name: "content",
						help: "contenu du post",
						next: Run(PostWrite),
					},
				},
			},
			{
				name: "reply",
				help: "répondre à un post",
				next: Select{
					name:    "topic",
					help:    "sujet de discussion",
					header:  "liste des sujets de discussions sur ce serveur",
					options: topicList,
					next: Select{
						name:    "post",
						help:    "message dans la discussion",
						header:  "liste des messages dans ce sujet de discussion",
						options: threadList,
						next: LongText{
							name: "content",
							help: "contenu de la réponse",
							next: Run(PostReply),
						},
					},
				},
			},
		},
	},
}

func PostRead(ctx Context) any {
	console := ctx.Value("console").(*Console)
	id := ctx.Value("post").(int)

	post, err := console.Server.Post(id, console.User)
	if err != nil {
		return ctx.Error(err)
	}

	b := strings.Builder{}

	fmt.Fprintf(&b, "ID : %d\n", id)
	if post.Parent != 0 {
		fmt.Fprintf(&b, "Réponse à : %d\n", post.Parent)
	}
	fmt.Fprintf(&b, "Auteur : %s\n", post.Author)
	fmt.Fprintf(&b, "Sujet : %s\n", post.Subject)
	fmt.Fprintln(&b, post.Content)

	return ctx.Output(b.String())
}

func PostWrite(ctx Context) any {
	console := ctx.Value("console").(*Console)
	subject := ctx.Value("subject").(string)
	content := ctx.Value("content").(string)

	post := Post{
		Server:  console.Server.Address,
		Date:    time.Now(),
		Author:  console.User.Login,
		Subject: subject,
		Content: content,
	}

	post, err := Save(post)
	if err != nil {
		return ctx.Error(err)
	}

	return ctx.Output(fmt.Sprintf("post %d ajouté au forum", post.ID))
}

func PostReply(ctx Context) any {
	console := ctx.Value("console").(*Console)
	id := ctx.Value("post").(int)
	content := ctx.Value("content").(string)

	original, err := console.Server.Post(id, console.User)
	if err != nil {
		return ctx.Error(err)
	}

	post := Post{
		Server:  console.Server.Address,
		Parent:  original.ID,
		Date:    time.Now(),
		Author:  console.User.Login,
		Subject: fmt.Sprintf("Re: %s", original.Subject),
		Content: content,
	}

	post, err = Save(post)
	if err != nil {
		return ctx.Error(err)
	}

	return ctx.Output(fmt.Sprintf("post %d ajouté au forum", post.ID))
}