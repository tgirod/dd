package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

type Client struct {
	width   int       // largeur de l'affichage
	height  int       // hauteur de l'affichage
	input   Input     // invite de commande
	output  string    // résultat de la dernière commande
	lastCmd string    // dernière commande saisie
	modal   tea.Model // interface modale

	Game    // état interne du jeu
	Console // console enregistrée dans le jeu
}

func (c Client) Init() tea.Cmd {
	return func() tea.Msg {
		// enregistrer la console dans l'état du jeu
		console, err := NewConsole(c.Game)
		if err != nil {
			return ErrorMsg{err}
		}
		return ConsoleMsg{console}
	}
}

// ErrorMsg contient le retour d'un programme à ajouter dans les logs
type ErrorMsg struct {
	Err error
}

func (e ErrorMsg) View(width int) string {
	return lg.PlaceHorizontal(width, lg.Center, errorTextStyle.Render(e.Err.Error()))
}

type ParseErrorMsg struct {
	Err  error
	Help string
}

func (p ParseErrorMsg) View(width int) string {
	b := strings.Builder{}
	b.WriteString(lg.PlaceHorizontal(width, lg.Center, errorTextStyle.Render(p.Err.Error())))
	b.WriteString(p.Help)
	return b.String()
}

// ConsoleMsg contient le nouvel état de la console
type ConsoleMsg struct {
	Console
}

// OpenModalMsg ouvre une fenêtre modale
type OpenModalMsg tea.Model

// CloseModalMsg ferme la fenêtre modale
type CloseModalMsg struct{}

func (c Client) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		c.height = msg.Height
		c.width = msg.Width
		return c, nil

	case ErrorMsg:
		c.output = msg.View(c.width)
		return c, nil

	case ParseErrorMsg:
		c.output = msg.View(c.width)
		return c, nil

	case HelpMsg:
		c.output = msg.Help

	case ConsoleMsg:
		c.Console = msg.Console

	case ConnectMsg:
		c.Console = msg.Console
		c.output = "connexion établie"
		return c, nil

	case JackMsg:
		c.Console = msg.Console
		c.output = "connexion illégale établie"
		if c.Console.Alarm == 0 {
			c.Console.Alarm = 1
		}
		return c, tea.Every(time.Second, c.Security)

	case LinkListMsg:
		c.output = msg.View()
		return c, nil

	case IndexMsg:
		c.output = msg.View()
		return c, nil

	case QuitMsg:
		c.output = "déconnexion"
		c.Server = Server{}
		c.Privilege = 0
		c.Alarm = 0
		c.Login = ""
		return c, c.Quit

	case DataSearchMsg:
		c.output = msg.View()
		return c, nil

	case OpenModalMsg:
		c.modal = msg
		c.input.Focus = false
		return c, nil

	case CloseModalMsg:
		c.modal = nil
		c.input.Focus = true
		return c, nil

	case SecurityScanMsg:
		return c, tea.Every(time.Second, c.Security)

	case SecurityIncreaseMsg:
		c.Alarm++
		return c, tea.Every(time.Second, c.Security)

	case SecurityKickMsg:
		c.output = "déconnecté de force du serveur"
		c.Server = Server{}
		c.Privilege = 0
		return c, c.Quit

	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			// quitter l'application client
			return c, tea.Sequentially(c.Quit, tea.Quit)
		}

		if c.modal != nil {
			// déléguer le traitement des messages à la fenêtre modale
			c.modal, cmd = c.modal.Update(msg)
			return c, cmd
		}

		if msg.Type == tea.KeyEnter {
			// valider la commande
			c.lastCmd = c.input.Value
			cmd = c.Run()
			c.input.Value = ""
			return c, cmd
		}

		// laisser le prompt gérer
		input, cmd := c.input.Update(msg)
		c.input = input.(Input)
		return c, cmd
	}

	return c, cmd
}

func (c Client) View() string {

	if c.modal == nil {
		return lg.JoinVertical(lg.Left,
			c.statusView(),
			c.outputView(),
			c.inputView(),
		)
	}

	modal := c.modal.View()

	return lg.Place(
		c.width, c.height,
		lg.Center, lg.Center,
		modal,
		lg.WithWhitespaceChars("+ "),
		lg.WithWhitespaceForeground(lg.Color("8")),
	)
}

var (
	// barre d'état
	statusStyle = lg.NewStyle().
			Padding(0, 1, 0, 1).
			Margin(0, 1, 0, 1).
			Foreground(lg.Color("0")).
			Background(lg.Color("10"))

	// affichage de la dernière commande
	outputStyle = lg.NewStyle().
			Padding(0, 1, 0, 1).
			Margin(0, 1, 0, 1)

	// invite de commande
	inputStyle = lg.NewStyle().
			Padding(0, 1, 0, 1).
			Margin(0, 1, 0, 1)

	// fenêtre modale
	modalStyle = lg.NewStyle().
			Padding(0, 1, 0, 1).
			Margin(0, 1, 0, 1).
			BorderStyle(lg.DoubleBorder()).
			BorderForeground(lg.Color("10"))

	focusFieldStyle = lg.NewStyle().
			BorderStyle(lg.NormalBorder()).
			BorderForeground(lg.Color("10"))

	unfocusFieldStyle = lg.NewStyle().
				BorderStyle(lg.NormalBorder()).
				BorderForeground(lg.Color("8"))

	// texte discret
	mutedTextStyle = lg.NewStyle().Foreground(lg.Color("8"))

	// texte normal
	normalTextStyle = lg.NewStyle().Foreground(lg.Color("15"))

	// curseur
	cursorStyle = lg.NewStyle().Reverse(true)

	// texte erreur
	errorTextStyle = lg.NewStyle().Foreground(lg.Color("9")).Padding(1)

	greenTextStyle  = lg.NewStyle().Foreground(lg.Color("0")).Background(lg.Color("10"))
	yellowTextStyle = lg.NewStyle().Foreground(lg.Color("0")).Background(lg.Color("11"))
	redTextStyle    = lg.NewStyle().Foreground(lg.Color("0")).Background(lg.Color("9"))
)

func (c Client) statusView() string {
	login := fmt.Sprintf("👤 %s", c.Console.Login)
	priv := strings.Repeat("✪", c.Console.Privilege)
	alarm := strings.Repeat("💀", c.Console.Alarm)

	width := c.width - statusStyle.GetHorizontalFrameSize() - lg.Width(alarm)
	status := lg.PlaceHorizontal(width, lg.Left, login+" "+priv) + alarm

	return statusStyle.Render(status)
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
	content = lg.Place(width, height, lg.Left, lg.Bottom, content)

	return outputStyle.Render(content)
}

func (c Client) inputView() string {
	content := c.input.View()
	width := c.width - inputStyle.GetHorizontalFrameSize()
	content = lg.PlaceHorizontal(width, lg.Left, "> "+content)
	return inputStyle.Render(content)
}

func NewClient(width, height int, game Game) Client {
	c := Client{
		width:  width,
		height: height,
		input: Input{
			Focus:       true,
			Placeholder: "help",
		},
		Game: game,
	}

	return c
}

// Run parse et exécute la commande saisie par l'utilisateur
func (c Client) Run() tea.Cmd {
	args := strings.Fields(c.input.Value)

	return func() tea.Msg {
		// construire la tea.Cmd qui parse et exécute la commande

		// exécuter la commande
		ctx := Context{c.Game, c.Console}
		return c.Console.Run(ctx, args)
	}
}

// Quit supprime la console de l'état du jeu
func (c Client) Quit() tea.Msg {
	if err := c.Game.DeleteStruct(&c.Console); err != nil {
		fmt.Println(err)
	}

	return nil
}

func (c Client) Security(t time.Time) tea.Msg {
	if c.Console.Alarm >= 10 {
		// hacker repéré, il se fait kicker du serveur
		return SecurityKickMsg{}
	}

	r := rand.Float64()
	if r < c.Console.Server.Detection {
		// la localisation du hacker progresse
		return SecurityIncreaseMsg{}
	}

	// la localisation du hacker se poursuit
	return SecurityScanMsg{}
}

type SecurityIncreaseMsg struct{}

type SecurityScanMsg struct{}

type SecurityKickMsg struct{}