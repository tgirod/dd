package main

import tea "github.com/charmbracelet/bubbletea"

type Link struct {
	service string // nom du service à suivre pour se connecter
}

func (c Link) Init() tea.Cmd {
	return nil
}

func (c Link) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return c, nil
}

func (c Link) View() string {
	return ""
}