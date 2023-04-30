package main

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type SelectKeymap struct {
	Select key.Binding
	Cancel key.Binding
}

var DefaultSelectKeymap = SelectKeymap{
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "sélectionner"),
	),
	Cancel: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "annuler"),
	),
}

func (k SelectKeymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Select, k.Cancel}
}

func (k SelectKeymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Select},
		{k.Cancel},
	}
}

func (o Option) FilterValue() string {
	return o.Filter
}

var selectedStyle = lg.NewStyle().Reverse(true)

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	item, ok := listItem.(Option)
	if !ok {
		return
	}

	str := item.Desc
	if index == m.Index() {
		str = selectedStyle.Render(str)
	}
	fmt.Fprintf(w, str)
}

var delegate = itemDelegate{}

type SelectModel struct {
	ctx  Context
	cmd  Cmd
	name string
	list list.Model
}

func NewSelect(ctx Context, cmd Cmd, name string, title string, options []Option) *SelectModel {
	items := make([]list.Item, len(options))
	for i, o := range options {
		items[i] = o
	}

	m := SelectModel{
		ctx:  ctx,
		cmd:  cmd,
		name: name,
		list: list.New(items, delegate, 0, 0),
	}

	m.list.DisableQuitKeybindings()

	return &m
}

func (m *SelectModel) Init() tea.Cmd {
	return nil
}

func (m *SelectModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultSelectKeymap.Select):
			return m.Select()
		case key.Matches(msg, DefaultSelectKeymap.Cancel):
			return m.Cancel()
		default:
			m.list, cmd = m.list.Update(msg)
			return m, cmd
		}
	default:
		m.list, cmd = m.list.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m *SelectModel) View() string {
	return m.list.View()
}

func (m *SelectModel) Select() (tea.Model, tea.Cmd) {
	// items sélectionné
	item := m.list.SelectedItem().(Option)

	// stocker le choix dans le contexte
	ctx := m.ctx.New(m.name, item.Value, m.cmd)

	cmd := tea.Batch(
		MsgToCmd(CloseModalMsg{}),
		MsgToCmd(ctx),
	)
	return m, cmd
}

func (m *SelectModel) Cancel() (tea.Model, tea.Cmd) {
	ctx := m.ctx.Cancel()
	return m, tea.Batch(
		MsgToCmd(CloseModalMsg{}),
		MsgToCmd(ctx),
	)
}
