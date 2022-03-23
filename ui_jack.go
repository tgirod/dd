package main

import tea "github.com/charmbracelet/bubbletea"

type JackConnect struct {
	address string
}

func (c JackConnect) Init() tea.Cmd {
	return nil
}

func (c JackConnect) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return c, nil
}

func (c JackConnect) View() string {
	return ""
}

type JackLink struct {
	service string // nom du service à suivre pour se connecter
}

func (c JackLink) Init() tea.Cmd {
	return nil
}

func (c JackLink) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return c, nil
}

func (c JackLink) View() string {
	return ""
}