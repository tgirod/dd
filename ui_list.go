package main

import (
	"strconv"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ListKeymap struct {
	Select key.Binding
	Cancel key.Binding
}

var DefaultListKeymap = ListKeymap{
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "sélectionner"),
	),
	Cancel: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "annuler"),
	),
}

func (k ListKeymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Select, k.Cancel}
}

func (k ListKeymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Select},
		{k.Cancel},
	}
}

type ListModel struct {
	ctx      Context
	list     list.Model
	delegate list.ItemDelegate // affichage d'un item
	cancel   bool              // la touche esc annule toute la commande
}

func NewList(ctx Context, items []list.Item, del list.ItemDelegate, cancel bool) *ListModel {
	if del == nil {
		del = list.NewDefaultDelegate()
	}

	m := ListModel{
		ctx:      ctx,
		list:     list.New(items, del, 0, 0),
		delegate: del,
		cancel:   cancel,
	}

	m.list.DisableQuitKeybindings()

	return &m
}

func (m *ListModel) Init() tea.Cmd {
	return nil
}

func (m *ListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultListKeymap.Select):
			return m.Select()
		case key.Matches(msg, DefaultListKeymap.Cancel):
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

func (m *ListModel) View() string {
	return m.list.View()
}

func (m *ListModel) Select() (tea.Model, tea.Cmd) {
	// récupérer l'index du curseur
	ctx := m.ctx
	index := m.list.Cursor()
	// ajouter l'index aux arguments et retourner le context
	ctx.Args = append(ctx.Args, strconv.Itoa(index))
	cmd := tea.Batch(
		MsgToCmd(CloseModalMsg{}),
		MsgToCmd(ctx),
	)
	return m, cmd
}

func (m *ListModel) Cancel() (tea.Model, tea.Cmd) {
	args := m.ctx.Args

	// annuler la commande
	if m.cancel || len(args) == 0 {
		return m, MsgToCmd(CloseModalMsg{})
	}

	// retirer le dernier argument et relancer la commande
	m.ctx.Args = args[0 : len(args)-1]
	cmd := tea.Batch(
		MsgToCmd(CloseModalMsg{}),
		MsgToCmd(m.ctx),
	)
	return m, cmd
}
