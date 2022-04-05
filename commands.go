package main

import tea "github.com/charmbracelet/bubbletea"

// Connect tente d'établir la connexion entre la console et le serveur à
// l'adresse indiquée
func (c Client) Connect(address, login, password string) tea.Msg {
	return LogMsg{msg: "connexion établie"}
}
