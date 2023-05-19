package main

import (
	"fmt"
	"strings"
	"time"

	lg "github.com/charmbracelet/lipgloss"
)

// Nouvelle organisation des Forums => à plat.
// Les Topics n'ont pas de parent
//   Un thread par Topic, on ne peut répondre qu'au Topic.
//   forum read (no #topic) => liste tous les Topics
//   forum read  #topic => liste tous les Post (avec content) du Topic
// TODO ajouter nb de réponse quand on liste les Topic ?

func topicList(ctx Context) ([]Option, error) {
	console := ctx.Value("console").(*Console)
	topics := console.Server.Topics(console.User)
	opts := make([]Option, 0, len(topics))
	for _, t := range topics {
		if allowedGroup(t.Group, console.User.Groups) {
			opts = append(opts, Option{
				help:  fmt.Sprintf("%s\t%s\t%s\t", t.Date.Format(time.DateTime), t.Author, t.Subject),
				value: t.ID,
			})
		}
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

// liste les différentes options de Groupe d'un User
func groupList(ctx Context) ([]Option, error) {
	console := ctx.Console()
	groups := console.User.Groups

	opts := []Option{
		{
			value: "public",
			help: "(tout le monde peut lire)",
		},
	}
	for _, g := range groups {
		opts = append(opts, Option{
			value: g,
			help: "réservé à ce groupe",
		})
	}
	return opts, nil
}
// check that group is either public or "" or in validGroups
func allowedGroup(group string, authorizedGroups []string) bool {
	if group == "" || group == "public" {
		return true
	}
	for _, v := range authorizedGroups {
		if group == v {
			return true
		}
	}

	return false
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
					next: Run(TopicRead),
					// next: Select{
					// 	name:    "post",
					// 	help:    "message dans la discussion",
					// 	header:  "liste des messages dans ce sujet de discussion",
					// 	options: threadList,
					// 	next:    Run(PostRead),
					// },
				},
			},
			{
				name: "write",
				help: "ouvrir un nouveau sujet",
				next: Select{
					name: "group",
					help:	"groupe propriétaire du sujet",
					header: "Sujet de discussion restreint au groupe: ",
					options: groupList,
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
			},
			{
				name: "answer",
				help: "répondre à un topic",
				next: Select{
					name:    "topic",
					help:    "sujet de discussion",
					header:  "liste des sujets de discussions sur ce serveur",
					options: topicList,
					// next: Select{
					// 	name:    "post",
					// 	help:    "message dans la discussion",
					// 	header:  "liste des messages dans ce sujet de discussion",
					// 	options: threadList,
						next: LongText{
							name: "content",
							help: "contenu de la réponse",
							next: Run(TopicAnswer),
							// },
					},
				},
			},
			{
				name: "dump",
				help: "dump all posts",
				next: Run(DumpForum),
			},
		},
	},
}

// style
var titleStyle = lg.NewStyle().Reverse(true)
// var contentStyle = lg.NewStyle().BorderStyle(lg.Border{Left: " | "}).BorderLeft(true)
var contentStyle = lg.NewStyle().MarginLeft(4)
func (t Thread) Render(prefix string) string {
	b := strings.Builder{}

	fmt.Fprintf(&b, "%s\n", titleStyle.Render(fmt.Sprintf("%3d %20s %15s %s %s",
		t.Post.ID, t.Date.Format(time.DateTime),
		t.Author, prefix, t.Subject)))
	fmt.Fprintf(&b, "%s", contentStyle.Render(t.Post.Content))

	return b.String()
}
// Tous les Posts (avec content) d'un Topic
func TopicRead(ctx Context) any {
	console := ctx.Value("console").(*Console)

	topic := ctx.Value("topic").(int)
	// récupérer le post racine
	root, err := console.Server.Post(topic, console.User)
	if err != nil {
		return ctx.Error(err)
	}
	// récupérer le thread : tous les posts
	t, err := console.Server.Thread(root, console.User)

	b := strings.Builder{}
	// D'abord affiche le Topic parent
	fmt.Fprintf(&b, "%s\n", t.Render(""))

	var prefix = ""
	for i, r := range t.Replies {
		if i < len(t.Replies)-1 {
			prefix = "├─ "
		} else {
			prefix = "└─ "
		}
		fmt.Fprintf(&b, "%s\n", r.Render(prefix))
	}

	return ctx.Output(b.String())
}
// func PostRead(ctx Context) any {
// 	console := ctx.Value("console").(*Console)
// 	id := ctx.Value("post").(int)

// 	post, err := console.Server.Post(id, console.User)
// 	if err != nil {
// 		return ctx.Error(err)
// 	}

// 	b := strings.Builder{}

// 	fmt.Fprintf(&b, "ID : %d\n", id)
// 	if post.Parent != 0 {
// 		fmt.Fprintf(&b, "Réponse à : %d\n", post.Parent)
// 	}
// 	fmt.Fprintf(&b, "Auteur : %s\n", post.Author)
// 	fmt.Fprintf(&b, "Sujet : %s\n", post.Subject)
// 	fmt.Fprintln(&b, post.Content)

// 	return ctx.Output(b.String())
// }

func PostWrite(ctx Context) any {
	console := ctx.Value("console").(*Console)
	group := ctx.Value("group").(string)
	if group == "public" {
		group = ""
	}
	subject := ctx.Value("subject").(string)
	content := ctx.Value("content").(string)

	post := Post{
		Server:  console.Server.Address,
		Group:	group,
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

func TopicAnswer(ctx Context) any {
	console := ctx.Value("console").(*Console)
	topic := ctx.Value("topic").(int)
	content := ctx.Value("content").(string)

	original, err := console.Server.Post(topic, console.User)
	if err != nil {
		return ctx.Error(err)
	}

	post := Post{
		Server:  console.Server.Address,
		Group:   original.Group,
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

// DEBUG
func DumpForum(ctx Context) any {
	console := ctx.Value("console").(*Console)
	SerializePosts(console.Server.Address)

	return ctx.Output("forum dumped on admin console")
}
