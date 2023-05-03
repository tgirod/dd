package main

import (
	"fmt"
	"strings"
	"time"
)

var forum = Cmd{
	Name:       "forum",
	ShortHelp:  "participer au forum du serveur",
	Connected:  true,
	Identified: true,
	Args:       []Arg{},
	SubCmds: []Cmd{
		{
			Name:      "read",
			ShortHelp: "lire le forum",
			Args: []Arg{
				{
					Type:      SelectNumberArg,
					Name:      "topic",
					ShortHelp: "sujet a lire",
					Options: func(ctx Context) []Option {
						console := ctx.Value("console").(*Console)
						posts := console.Server.Posts
						opts := make([]Option, 0, len(posts))
						for i, p := range posts {
							// un post est un topic si il est son propre parent
							if p.Parent == i {
								opts = append(opts, Option{
									Desc:  fmt.Sprintf("%d -- %s -- %s", i, p.Author, p.Subject),
									Value: i,
								})
							}
						}
						return opts
					},
				},
				{
					Type:      SelectNumberArg,
					Name:      "post",
					ShortHelp: "post a lire",
					Options: func(ctx Context) []Option {
						console := ctx.Value("console").(*Console)
						topic := ctx.Value("topic").(int)
						posts := console.Server.Posts
						opts := make([]Option, 0, len(posts))
						for i, p := range posts {
							if p.Parent == topic {
								opts = append(opts, Option{
									Desc:  fmt.Sprintf("%d -- %s -- %s", i, p.Author, p.Subject),
									Value: i,
								})
							}
						}
						return opts
					},
				},
			},
			Run: PostRead,
		},
		{
			Name:      "write",
			ShortHelp: "écrire un post",
			Args: []Arg{
				{
					Type:      ShortArg,
					Name:      "subject",
					ShortHelp: "sujet du post",
				},
				{
					Type:      LongArg,
					Name:      "content",
					ShortHelp: "contenu du post",
				},
			},
			SubCmds: []Cmd{},
			Run:     PostWrite,
		},
		{
			Name:      "reply",
			ShortHelp: "répondre à un message",
			Args: []Arg{
				{
					Type:      SelectNumberArg,
					Name:      "topic",
					ShortHelp: "sujet de discussion",
					Options: func(ctx Context) []Option {
						console := ctx.Value("console").(*Console)
						posts := console.Server.Posts
						opts := make([]Option, 0, len(posts))
						for i, p := range posts {
							// un post est un topic si il est son propre parent
							if p.Parent == i {
								opts = append(opts, Option{
									Desc:  fmt.Sprintf("%d -- %s -- %s", i, p.Author, p.Subject),
									Value: i,
								})
							}
						}
						return opts
					},
				},
				{
					Type:      SelectNumberArg,
					Name:      "post",
					ShortHelp: "post auquel répondre",
					Options: func(ctx Context) []Option {
						console := ctx.Value("console").(*Console)
						topic := ctx.Value("topic").(int)
						posts := console.Server.Posts
						opts := make([]Option, 0, len(posts))
						for i, p := range posts {
							if p.Parent == topic {
								opts = append(opts, Option{
									Desc:  fmt.Sprintf("%d -- %s -- %s", i, p.Author, p.Subject),
									Value: i,
								})
							}
						}
						return opts
					},
				},
				{
					Type:      LongArg,
					Name:      "content",
					ShortHelp: "contenu du post",
				},
			},
			Run: PostReply,
		},
	},
}

func PostRead(ctx Context) any {
	console := ctx.Value("console").(*Console)
	id := ctx.Value("post").(int)
	res := ctx.Result()

	if id < 0 || id >= len(console.Server.Posts) {
		res.Error = errInvalidArgument
		return res
	}

	post := console.Server.Posts[id]

	b := strings.Builder{}

	fmt.Fprintf(&b, "ID : %d\n", id)
	if post.Parent != 0 {
		fmt.Fprintf(&b, "Réponse à : %d\n", post.Parent)
	}
	fmt.Fprintf(&b, "Auteur : %s\n", post.Author)
	fmt.Fprintf(&b, "Sujet : %s\n", post.Subject)
	fmt.Fprintln(&b, post.Content)

	res.Output = b.String()
	return res
}

func PostWrite(ctx Context) any {
	console := ctx.Value("console").(*Console)
	subject := ctx.Value("subject").(string)
	content := ctx.Value("content").(string)
	res := ctx.Result()

	post := Post{
		Parent:  len(console.Posts),
		Date:    time.Now(),
		Author:  console.Account.Login,
		Subject: subject,
		Content: content,
	}

	console.Server.Posts = append(console.Server.Posts, post)

	res.Output = fmt.Sprintf("post %d ajouté au forum", len(console.Server.Posts))
	return res
}

func PostReply(ctx Context) any {
	console := ctx.Value("console").(*Console)
	id := ctx.Value("post").(int)
	content := ctx.Value("content").(string)
	res := ctx.Result()

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

	res.Output = fmt.Sprintf("post %d ajouté au forum", len(console.Server.Posts))
	return res
}