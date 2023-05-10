package main

import (
	"fmt"
	"strings"

	"github.com/asdine/storm/v3/q"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mattn/go-shellwords"
)

type Admin struct {
	prompt textinput.Model
}

func (a Admin) Init() tea.Cmd {
	return textinput.Blink
}

func (a Admin) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return a, tea.Quit
		case tea.KeyEnter:
			prompt := a.prompt.Value()
			a.prompt.Reset()
			return a, a.Parse(prompt)
		default:
			a.prompt, cmd = a.prompt.Update(msg)
			return a, cmd
		}
	default:
		a.prompt, cmd = a.prompt.Update(msg)
		return a, cmd
	}
}

func (a Admin) View() string {
	return a.prompt.View()
}

func NewAdmin() Admin {
	prompt := textinput.New()
	prompt.Focus()
	return Admin{
		prompt: prompt,
	}
}

func (a Admin) Parse(prompt string) tea.Cmd {
	args, err := shellwords.Parse(prompt)
	if err != nil {
		return tea.Println("erreur", err)
	}

	ctx := Context{
		parent: &Context{},
		key:    "console",
		value:  nil,
		node:   admin,
	}

	switch res := admin.Parse(ctx, args).(type) {
	case tea.Cmd:
		return res
	case Result:
		return tea.Println(res.String())
	default:
		return tea.Println("bizarre")
	}
}

func AdminStart() {
	p := tea.NewProgram(NewAdmin())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

var admin = Branch{
	name: "",
	cmds: []Cmd{
		{
			name: "id",
			help: "manipuler les identités",
			next: Branch{
				name: "",
				cmds: []Cmd{
					{
						name: "list",
						help: "lister les identités",
						next: Run(ListIdentity),
					},
				},
			},
		},
		{
			name: "msg",
			help: "manipuler les messages",
			next: Branch{
				name: "",
				cmds: []Cmd{
					{
						name: "from",
						help: "messages envoyés par un utilisateur",
						next: Select{
							name:   "from",
							help:   "auteur des messages",
							header: "liste des auteurs",
							options: func(ctx Context) ([]Option, error) {
								identities, err := Identities()
								if err != nil {
									return []Option{}, err
								}
								opts := make([]Option, len(identities))
								for i, identity := range identities {
									opts[i] = Option{
										value: identity.Login,
										help:  identity.Name,
									}
								}
								return opts, nil
							},
							next: Run(MsgFrom),
						},
					},
					{
						name: "to",
						help: "messages envoyés à un utilisateur",
						next: Select{
							name:   "to",
							help:   "destinataire des messages",
							header: "liste des destinataires",
							options: func(ctx Context) ([]Option, error) {
								identities, err := Identities()
								if err != nil {
									return []Option{}, err
								}
								opts := make([]Option, len(identities))
								for i, identity := range identities {
									opts[i] = Option{
										value: identity.Login,
										help:  identity.Name,
									}
								}
								return opts, nil
							},
							next: Run(MsgTo),
						},
					},
				},
			},
		},
	},
}

// ListIdentity liste les identités
func ListIdentity(ctx Context) any {
	identities, err := Identities()
	if err != nil {
		return ctx.Error(err)
	}

	b := strings.Builder{}
	for _, i := range identities {
		fmt.Fprintf(&b, "%+v\n", i)
	}

	return tea.Println(b.String())
}

func MsgFrom(ctx Context) any {
	from := ctx.Value("from").(string)
	messages, err := Find[Message](
		q.Eq("From", from),
	)
	if err != nil {
		return tea.Println(err)
	}

	b := strings.Builder{}
	fmt.Fprintf(&b, "messages envoyés par %s : %d\n", from, len(messages))

	for _, m := range messages {
		fmt.Fprintln(&b, m.String())
	}

	return tea.Println(b.String())
}

func MsgTo(ctx Context) any {
	to := ctx.Value("to").(string)
	messages, err := Find[Message](
		q.Eq("To", to),
	)
	if err != nil {
		return tea.Println(err)
	}

	b := strings.Builder{}
	fmt.Fprintf(&b, "messages envoyés à %s : %d\n", to, len(messages))

	for _, m := range messages {
		fmt.Fprintln(&b, m.String())
	}

	return tea.Println(b.String())
}
