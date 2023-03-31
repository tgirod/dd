package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/wordwrap"
)

type Pop struct {
	msg           string
	width, height int
}

func (p *Pop) Init() tea.Cmd {
	return nil
}

func (p *Pop) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		p.width = msg.Width
		p.height = msg.Height
	case tea.KeyMsg:
		return p, func() tea.Msg { return CloseModalMsg{} }
	}
	return p, nil
}

func (p *Pop) View() string {
	return wordwrap.String(p.msg, p.width)
}

var pop = Cmd{
	Name:      "pop",
	ShortHelp: "ouvre une fenêtre modale",
	Parse: func(args []string) any {
		msg := strings.Join(args, " ")
		if len(args) == 0 {
			msg = "ceci est une fenêtre modale"
		}
		return OpenModalMsg(&Pop{msg: msg})
	},
}
