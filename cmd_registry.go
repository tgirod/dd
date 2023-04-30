package main

import (
	"fmt"
	"strings"
)

var registry = Cmd{
	Name:      "registry",
	ShortHelp: "liste et manipule les registres du serveur",
	Connected: true,
	SubCmds: []Cmd{
		{
			Name:      "search",
			ShortHelp: "recherche dans les registres",
			Args: []Arg{
				{
					Name:      "prefix",
					ShortHelp: "préfixe du nom du registre",
					Type:      ShortArg,
				},
			},
			Run: RegistrySearch,
		},
		{
			Name:      "edit",
			ShortHelp: "modifie un registre",
			Args: []Arg{
				{
					Name:      "name",
					ShortHelp: "nom du registre à modifier",
					Type:      ShortArg,
				},
			},
			Run: RegistryEdit,
		},
	},
}

func RegistryEdit(ctx Context) any {
	console := ctx.Value("console").(*Console)
	name := ctx.Value("name").(string)
	res := ctx.Result()

	state, err := console.Server.RegistryEdit(name)

	if err != nil {
		res.Error = err
		return res
	}

	res.Output = fmt.Sprintf("nouvel état du registre %s : %v\n", name, state)
	return res
}

func RegistrySearch(ctx Context) any {
	console := ctx.Value("console").(*Console)
	prefix := ctx.Value("prefix").(string)
	res := ctx.Result()

	search := console.Server.RegistrySearch(prefix)

	b := strings.Builder{}
	tw := tw(&b)
	fmt.Fprintf(tw, "NAME\tSTATE\tDESCRIPTION\t\n")
	for _, r := range search {
		fmt.Fprintf(tw, "%s\t%t\t%s\t\n", r.Name, r.State, r.Description)
	}
	tw.Flush()

	res.Output = b.String()
	return res
}