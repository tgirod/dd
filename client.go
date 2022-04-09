package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

type Client struct {
	width   int             // largeur de l'affichage
	height  int             // hauteur de l'affichage
	input   textinput.Model // invite de commande
	output  string          // résultat de la dernière commande
	lastCmd string          // dernière commande saisie

	Game    // état interne du jeu
	Console // console enregistrée dans le jeu
}

func (c Client) Init() tea.Cmd {
	return func() tea.Msg {
		// enregistrer la console dans l'état du jeu
		console, err := NewConsole(c.Game)
		if err != nil {
			return LogMsg{err: err}
		}
		return ConsoleMsg{console}
	}
}

// ConsoleMsg contient le nouvel état de la console
type ConsoleMsg struct {
	Console
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
			c.lastCmd = c.input.Value()
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
		// affiche le résultat de la commande dans la sortie
		c.output = msg.View()
		return c, nil

	case HelpMsg:
		// affiche l'aide dans la sortie
		c.output = msg.Help

	case ConsoleMsg:
		// mettre à jour la console associée au client
		c.Console = msg.Console

	case ConnectMsg:
		c.Console = msg.Console
		c.output = "connexion établie"
	}

	c.input, cmd = c.input.Update(msg)
	return c, cmd
}

func (c Client) View() string {
	return lg.JoinVertical(lg.Left,
		c.statusView(),
		c.outputView(),
		c.inputView(),
	)
}

var (
	// barre d'état
	statusStyle = lg.NewStyle().
			Padding(0, 1, 0, 1).
			Margin(0, 1, 0, 1).
			Background(lg.Color("2")).
			Foreground(lg.Color("15"))

	// affichage de la dernière commande
	outputStyle = lg.NewStyle().
			PaddingLeft(1).PaddingRight(1).
			Margin(0, 1, 0, 1).
			BorderStyle(lg.NormalBorder()).
			BorderForeground(lg.Color("10"))

	// invite de commande
	inputStyle = lg.NewStyle().
			PaddingLeft(1).PaddingRight(1)
)

func (c Client) statusView() string {
	status := fmt.Sprintf("privilege: %d", c.Console.Privilege)
	width := c.width - statusStyle.GetHorizontalFrameSize()
	return statusStyle.Render(
		lg.PlaceHorizontal(width, lg.Left, status),
	)
}

func (c Client) outputView() string {
	// dimensions de l'espace d'affichage
	width := c.width - outputStyle.GetHorizontalFrameSize()
	height := c.height - 2 - outputStyle.GetVerticalFrameSize()

	// dernière commande + output
	content := ""
	if c.lastCmd != "" {
		content = lg.JoinVertical(lg.Left,
			fmt.Sprintf("> %s\n", c.lastCmd),
			c.output,
		)
	} else {
		content = c.output
	}

	// wrap au cas ou certaines lignes seraient trop longues
	content = wordwrap.String(content, width)

	// disposer le texte dans un espace qui remplit l'écran
	content = lg.Place(width, height, lg.Left, lg.Top, content)

	return outputStyle.Render(content)
}

func (c Client) inputView() string {
	width := c.width - inputStyle.GetHorizontalFrameSize()
	content := lg.PlaceHorizontal(width, lg.Left, c.input.View())
	return inputStyle.Render(content)
}

func NewClient(width, height int, game Game) Client {
	c := Client{
		width:  width,
		height: height,
		input:  textinput.New(),
		Game:   game,
	}
	c.input.Focus()
	c.input.Placeholder = "entrez une commande ou help"

	return c
}

// Run parse et exécute la commande saisie par l'utilisateur
func (c Client) Run() tea.Cmd {
	args := strings.Fields(c.input.Value())

	return func() tea.Msg {
		// construire la tea.Cmd qui parse et exécute la commande

		// exécuter la commande
		ctx := Context{c.Game, c.Console}
		return c.Console.Run(ctx, args)
	}
}

// Quit supprime la console de l'état du jeu
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
