package main

import (
	"fmt"
	"strings"
)

var registry = Cmd{
	name:      "registry",
	help:      "liste et manipule les registres du serveur",
	connected: true,
	next: Branch{
		name: "",
		cmds: []Cmd{
			{
				name: "search",
				help: "recherche dans les registres",
				next: String{
					name: "prefix",
					help: "préfixe du nom du registre",
					next: Run(RegistrySearch),
				},
			},
			{
				name: "edit",
				help: "modifie un registre",
				next: String{
					name: "name",
					help: "nom du registre à modifier",
					next: Run(RegistryEdit),
				},
			},
		},
	},
}

func RegistryEdit(ctx Context) any {
	console := ctx.Value("console").(*Console)
	name := ctx.Value("name").(string)

	state, err := console.Server.RegistryEdit(name)

	if err != nil {
		return ctx.Error(err)
	}

	return ctx.Output(fmt.Sprintf("nouvel état du registre %s : %v\n", name, state))
}

func RegistrySearch(ctx Context) any {
	console := ctx.Value("console").(*Console)
	prefix := ctx.Value("prefix").(string)

	search := console.Server.RegistrySearch(prefix)

	b := strings.Builder{}
	tw := tw(&b)
	fmt.Fprintf(tw, "NAME\tSTATE\tDESCRIPTION\t\n")
	for _, r := range search {
		fmt.Fprintf(tw, "%s\t%t\t%s\t\n", r.Name, r.State, r.Description)
	}
	tw.Flush()

	return ctx.Output(b.String())
}