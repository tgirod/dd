package main

import (
	"fmt"
	"strings"
	"time"
)

func topicList(ctx Context) []Option {
	console := ctx.Value("console").(*Console)
	posts := console.Server.Posts
	opts := make([]Option, 0, len(posts))
	for i, p := range posts {
		// un post est un topic si il est son propre parent
		if p.Parent == i {
			opts = append(opts, Option{
				help:  fmt.Sprintf("%s -- %s", p.Author, p.Subject),
				value: i,
			})
		}
	}
	return opts
}

func postList(ctx Context) []Option {
	console := ctx.Value("console").(*Console)
	topic := ctx.Value("topic").(int)
	posts := console.Server.Posts
	opts := make([]Option, 0, len(posts))
	for i, p := range posts {
		if p.Parent == topic {
			opts = append(opts, Option{
				help:  fmt.Sprintf("%s -- %s", p.Author, p.Subject),
				value: i,
			})
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
						options: postList,
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
						options: postList,
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

	post := console.Server.Posts[id]

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
		Parent:  len(console.Posts),
		Date:    time.Now(),
		Author:  console.Account.Login,
		Subject: subject,
		Content: content,
	}

	console.Server.Posts = append(console.Server.Posts, post)

	return ctx.Output(fmt.Sprintf("post %d ajouté au forum", len(console.Server.Posts)))
}

func PostReply(ctx Context) any {
	console := ctx.Value("console").(*Console)
	id := ctx.Value("post").(int)
	content := ctx.Value("content").(string)

	original := console.Posts[id]
	parent := console.Posts[original.Parent]

	post := Post{
		Parent:  original.Parent,
		Date:    time.Now(),
		Author:  console.Account.Login,
		Subject: fmt.Sprintf("Re: %s", parent.Subject),
		Content: content,
	}

	console.Server.Posts = append(console.Server.Posts, post)

	return ctx.Output(fmt.Sprintf("post %d ajouté au forum", len(console.Server.Posts)))
}