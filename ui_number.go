package main

import (
	"strconv"

	hhelp "github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type NumberKeymap struct {
	Validate key.Binding
	Cancel   key.Binding
}

var DefaultNumberKeymap = NumberKeymap{
	Validate: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "valider"),
	),
	Cancel: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "annuler"),
	),
}

func (k NumberKeymap) ShortHelp() []key.Binding {
	return []key.Binding{k.Validate, k.Cancel}
}

func (k NumberKeymap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Validate},
		{k.Cancel},
	}
}

type NumberModel struct {
	ctx   Context // contexte a exécuter après saisie
	cmd   Cmd     // commande pour laquelle la fenêtre est ouverte
	name  string  // utilisé comme clef pour stocker le résultat
	title string  // titre de la fenêtre modale
	input textinput.Model
	help  hhelp.Model
}

func NewNumber(ctx Context, cmd Cmd, name string, title string) *NumberModel {
	m := NumberModel{
		ctx:   ctx,
		cmd:   cmd,
		name:  name,
		title: title,
		input: textinput.New(),
		help:  hhelp.New(),
	}

	m.input.Placeholder = name
	m.input.Width = 30

	m.help.Width = UI_WIDTH

	return &m
}

func (m *NumberModel) Init() tea.Cmd {
	return m.input.Focus()
}

func (m *NumberModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, DefaultNumberKeymap.Validate):
			return m.Validate()
		case key.Matches(msg, DefaultNumberKeymap.Cancel):
			return m.Cancel()
		default:
			m.input, cmd = m.input.Update(msg)
		}
	default:
		m.input, cmd = m.input.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m *NumberModel) View() string {
	title := uiStyle.Copy().Align(lg.Center).MarginBottom(1).Render(m.title)
	input := m.input.View()
	help := uiStyle.Copy().MarginTop(1).Render(m.help.View(DefaultNumberKeymap))
	return uiStyle.Render(lg.JoinVertical(lg.Left, title, input, help))
}

// Validate ajoute la saisie au contexte et relance l'exécution
func (m *NumberModel) Validate() (tea.Model, tea.Cmd) {
	value, err := strconv.Atoi(m.input.Value())
	if err != nil {
		m.input.Reset()
		return m, nil
	}

	// stocker dans le contexte
	ctx := m.ctx.New(m.name, value, m.cmd)

	// retourner le contexte à relancer
	cmd := tea.Batch(
		MsgToCmd(CloseModalMsg{}),
		MsgToCmd(ctx),
	)
	return m, cmd
}

func (m *NumberModel) Cancel() (tea.Model, tea.Cmd) {
	ctx := m.ctx.Cancel()
	return m, tea.Batch(
		MsgToCmd(CloseModalMsg{}),
		MsgToCmd(ctx),
	)
}
