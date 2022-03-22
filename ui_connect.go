package main

import tea "github.com/charmbracelet/bubbletea"

// Connect est le programme permettant d'établir la connexion a un
// serveur distant.
type Connect struct {
	Address  string // adresse du serveur à joindre
	Login    string // identifiant d'accès
	Password string // mot de passe
}

func (c Connect) Init() tea.Cmd {
	return nil
}

func (c Connect) Update(tea.Msg) (tea.Model, tea.Cmd) {
	return c, nil
}

func (c Connect) View() string {
	return ""
}