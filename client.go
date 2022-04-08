package main

import (
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func NoOp() tea.Msg {
	return nil
}

type Client struct {
	width  int    // largeur de l'affichage
	height int    // hauteur de l'affichage
	input  string // saisie utilisateur
	log    string // résultat de la dernière commande

	game      Game // état interne du jeu
	consoleID int  // identifiant de la console
}

func (c Client) Init() tea.Cmd {
	return nil
}

func (c Client) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c":
			// quitter l'application client
			return c, tea.Sequentially(c.Quit, tea.Quit)

		case "enter":
			// lancer l'exécution de la commande
			cmd := c.Run()
			c.input = ""
			return c, cmd
		}

		if msg.Type == tea.KeyRunes {
			// ajouter dans le champ input
			c.input += msg.String()
		}

		if msg.Type == tea.KeyBackspace {
			if len(c.input) == 0 {
				return c, nil
			}

			// supprimer la dernière rune
			input := []rune(c.input)
			if len(input) > 0 {
				c.input = string(input[:len(input)-1])
			}
		}

	case tea.WindowSizeMsg:
		// gère le redimensionnement de la fenêtre
		c.height = msg.Height
		c.width = msg.Width
		return c, nil

	case LogMsg:
		// ajoute dans les logs
		c.log = msg.View()
		return c, nil
	}

	return c, nil
}

func (c Client) View() string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "%s\n", c.log)
	fmt.Fprintf(&b, "> %s\n", c.input)
	return b.String()
}

func NewClient(width, height int, game Game) Client {
	co := NewConsole()
	if err := game.Save(&co); err != nil {
		log.Panic(err)
	}

	c := Client{
		width:     width,
		height:    height,
		game:      game,
		consoleID: co.ID,
	}

	return c
}

// Run parse et exécute la commande saisie par l'utilisateur
func (c Client) Run() tea.Cmd {
	args := strings.Fields(c.input)

	return func() tea.Msg {
		fmt.Println("run", args)
		// construire la tea.Cmd qui parse et exécute la commande

		// récupérer la console
		var console Console
		if err := c.game.One("ID", c.consoleID, &console); err != nil {
			return LogMsg{
				err: err,
			}
		}

		// exécuter la commande
		ctx := Context{c.game, console}
		return console.Run(ctx, args)
	}
}

func (c Client) Quit() tea.Msg {
	if err := c.game.Delete("Console", c.consoleID); err != nil {
		fmt.Println(err)
	}

	return nil
}

// LogMsg contient le retour d'un programme à ajouter dans les logs
type LogMsg struct {
	err error
	msg string
}

func (l LogMsg) View() string {
	b := strings.Builder{}
	if l.err != nil {
		fmt.Fprintf(&b, "ERR : %s\n\n", l.err.Error())
	}
	fmt.Fprintf(&b, "%s\n", l.msg)
	return b.String()
}
