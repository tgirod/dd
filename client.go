package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Client struct {
	width  int             // largeur de l'affichage
	height int             // hauteur de l'affichage
	input  textinput.Model // invite de commande
	output string          // résultat de la dernière commande

	Game    // état interne du jeu
	Console // console enregistrée dans le jeu
}

func (c Client) Init() tea.Cmd {
	return func() tea.Msg {
		// enregistrer la console dans l'état du jeu
		console, err := c.Game.NewConsole()
		if err != nil {
			return LogMsg{err: err}
		}
		return console
	}
}

func (c Client) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.Type {

		case tea.KeyCtrlC:
			// quitter l'application client
			return c, tea.Sequentially(c.Quit, tea.Quit)

		case tea.KeyEnter:
			// lancer l'exécution de la commande
			cmd = c.Run()
			c.input.Reset()
			return c, cmd
		}

	case tea.WindowSizeMsg:
		// gère le redimensionnement de la fenêtre
		c.height = msg.Height
		c.width = msg.Width
		return c, nil

	case LogMsg:
		// ajoute dans les logs
		c.output = msg.View()
		return c, nil

	case Console:
		// mettre à jour la console associée au client
		fmt.Printf("%+v\n", msg)
		c.Console = msg

	}

	c.input, cmd = c.input.Update(msg)
	return c, cmd
}

func (c Client) View() string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "%s\n", c.output)
	fmt.Fprint(&b, c.input.View())
	return b.String()
}

func NewClient(width, height int, game Game) Client {
	c := Client{
		width:  width,
		height: height,
		input:  textinput.New(),
		Game:   game,
	}
	c.input.Focus()

	return c
}

// Run parse et exécute la commande saisie par l'utilisateur
func (c Client) Run() tea.Cmd {
	args := strings.Fields(c.input.Value())

	return func() tea.Msg {
		fmt.Println("run", args)
		// construire la tea.Cmd qui parse et exécute la commande

		// exécuter la commande
		ctx := Context{c.Game, c.Console}
		return c.Console.Run(ctx, args)
	}
}

func (c Client) Quit() tea.Msg {
	if err := c.Game.DeleteStruct(c.Console); err != nil {
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
