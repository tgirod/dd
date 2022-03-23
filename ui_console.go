package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/davecgh/go-spew/spew"
)

type Console struct {
	network *Network    // données partagées du jeu
	width   int         // largeur de l'affichage
	height  int         // hauteur de l'affichage
	stack   []tea.Model // pile des vues ouvertes pour la navigation
}

func (c Console) Init() tea.Cmd {
	return func() tea.Msg {
		return NewPrompt()
	}
}

type Close struct{}

func (c Console) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if model, ok := msg.(tea.Model); ok {
		c.stack = append(c.stack, model)
		model.Init()
		return c, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// les bindings communs à toute l'appli
		spew.Dump(msg.String())
		switch msg.String() {
		case "ctrl+c":
			return c, tea.Quit
		}

	case tea.WindowSizeMsg:
		// gère le redimensionnement de la fenêtre
		c.height = msg.Height
		c.width = msg.Width
		return c, nil

	case Close:
		// ferme la dernière vue ouverte
		n := len(c.stack)
		if n > 0 {
			c.stack = c.stack[:n-1]
		}
		return c, nil
	}

	// délègue update à la vue au sommet de la pile
	top := len(c.stack) - 1
	m := c.stack[top]
	m, cmd := m.Update(msg)
	c.stack[top] = m
	return c, cmd
}

// View délègue l'affichage à la vue au sommet de la pile
func (c Console) View() string {
	if len(c.stack) == 0 {
		return ""
	}
	return c.stack[len(c.stack)-1].View()
}
