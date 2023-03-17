package main

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

type Pop struct {
	msg           string
	width, height int
}

func (p Pop) ParseName() string {
	return "pop"
}

func (p Pop) ShortHelp() string {
	return "ouvre un popup"
}

func (p Pop) LongHelp() string {
	b := strings.Builder{}
	b.WriteString(p.ShortHelp() + "\n")
	b.WriteString("\nUSAGE\n")
	b.WriteString("  pop <MSG>\n")
	b.WriteString("\nARGUMENTS\n")
	b.WriteString("  <MSG> -- message à afficher")
	return b.String()
}

func (p Pop) Run(c *Client, args []string) tea.Msg {
	msg := strings.Join(args, " ")
	if len(args) == 0 {
		msg = "ceci est une fenêtre modale"
	}
	return OpenModalMsg(&Pop{msg, c.width, c.height})
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

var border = lg.NewStyle().Border(lg.DoubleBorder()).Padding(1)

func (p *Pop) View() string {
	content := border.Render(wordwrap.String(p.msg, p.width-border.GetWidth()))
	return lg.Place(p.width, p.height, lg.Center, lg.Center, content, lg.WithWhitespaceChars(". "))
}