package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/gliderlabs/ssh"
)

type Console struct {
	network *Network
	width   int
	height  int
	prompt  textinput.Model
	result  string
}

// NewConsole créé le modèle bubbletea qui sera utilisé pour l'affichage
func NewConsole(pty ssh.Pty) Console {
	c := Console{
		width:  pty.Window.Width,
		height: pty.Window.Height,
	}
	// initialiser l'invite de commande
	ti := textinput.New()
	ti.Placeholder = "commande"
	ti.Focus()
	ti.Width = c.width
	c.prompt = ti
	return c
}

// TODO Init initialise le modèle
func (c Console) Init() tea.Cmd {
	return nil
}

func (c Console) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// gestion du prompt
	if c.prompt, cmd = c.prompt.Update(msg); cmd != nil {
		return c, cmd
	}

	switch msg := msg.(type) {

	case resultMsg:
		// affiche le résultat de la commande à l'écran
		c.result = string(msg)
		c.prompt.Reset()
		return c, nil

	case tea.WindowSizeMsg:
		// gère le redimensionnement de la fenêtre
		c.height = msg.Height
		c.width = msg.Width

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			// Ctrl-C ferme l'application
			return c, tea.Quit
		case tea.KeyEnter:
			// Entrée lance le parsing de la commande
			return c, c.parseCommand(c.prompt.Value())
		default:
			c.prompt, cmd = c.prompt.Update(msg)
		}

	default:
		c.prompt, cmd = c.prompt.Update(msg)
	}

	return c, cmd
}

func (c Console) View() string {
	return c.result + "\n" + c.prompt.View()
}

type resultMsg string

func (c Console) parseCommand(input string) tea.Cmd {
	return func() tea.Msg {
		return resultMsg(input + " OK")
	}
}
