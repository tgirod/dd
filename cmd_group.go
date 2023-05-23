package main

var group = Cmd{
	name:       "group",
	help:       "gérer les groupes d'utilisateurs",
	connected:  true,
	identified: true,
	next: Branch{
		name: "",
		cmds: []Cmd{
			{
				name: "new",
				help: "créer un nouveau groupe",
				next: Text{
					name: "name",
					help: "nom du groupe",
					next: Run(GroupNew),
				},
			},
			{
				name: "invite",
				help: "inviter un utilisateur à rejoindre un groupe",
				next: Select{
					name:   "group",
					help:   "nom du groupe",
					header: "liste des groupes dont vous êtes administrateur",
					options: func(ctx Context) ([]Option, error) {
						c := ctx.Console()
						groups := c.GroupAdmin()
						opts := make([]Option, len(groups))
						for i, g := range groups {
							opts[i].help = g.Name
							opts[i].value = g.ID
						}
						return opts, nil
					},
					next: Text{
						name: "user",
						help: "nom de l'utilisateur",
						next: Run(GroupInvite),
					},
				},
			},
			{
				name: "ban",
				help: "bannir un utilisateur d'un groupe",
				next: Select{
					name:   "group",
					help:   "nom du groupe",
					header: "liste des groupes dont vous êtes administrateur",
					options: func(ctx Context) ([]Option, error) {
						c := ctx.Console()
						groups := c.GroupAdmin()
						opts := make([]Option, len(groups))
						for i, g := range groups {
							opts[i].value = g.Name
						}
						return opts, nil
					},
					next: Select{
						name: "user",
						help: "nom de l'utilisateur",
						options: func(ctx Context) ([]Option, error) {
							c := ctx.Console()
							group := ctx.Value("group").(string)
							members := c.Members(group)
							opts := make([]Option, len(members))
							for i, m := range members {
								opts[i].value = m.User
							}
							return opts, nil
						},
						next: Run(GroupBan),
					},
				},
			},
		},
	},
}

func GroupNew(ctx Context) any {
	return nil
}

func GroupInvite(ctx Context) any {
	return nil
}

func GroupBan(ctx Context) any {
	return nil
}

// group new <name>
// group invite <name> <user>
// group ban <name> <user>
