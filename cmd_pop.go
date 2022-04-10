package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type Pop struct{}

func (p Pop) ParseName() string {
	return "pop"
}

func (p Pop) ShortHelp() string {
	return "pop -- ouvre un popup"
}

func (p Pop) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(p.ShortHelp() + "\n")
	b.WriteString("USAGE\n")
	b.WriteString("  pop\n")
	return b.String()
}

func (p Pop) Run(ctx Context, args []string) tea.Msg {
	return OpenModalMsg(p)
}

func (p Pop) Init() tea.Cmd {
	return nil
}

func (p Pop) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return p, func() tea.Msg { return CloseModalMsg{} }
	}

	return p, nil
}

func (p Pop) View() string {
	pop := lg.Place(20, 3, lg.Center, lg.Center, "UNE PAGE DE PUB")
	return modalStyle.Render(pop)
}
