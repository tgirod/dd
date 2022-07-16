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

	*Game    // état interne du jeu
	*Console // console enregistrée dans le jeu
}

func NewClient(width, height int, game *Game) *Client {
	return &Client{
		width:  width,
		height: height,
		input: Input{
			Focus:       true,
			Placeholder: "help",
		},
		Game:    game,
		Console: NewConsole(),
	}
}

func (c *Client) Init() tea.Cmd {
	return nil
}

// affiche le résultat d'une commande
type ResultMsg struct {
	Error  error
	Output string
}

// OpenModalMsg ouvre une fenêtre modale
type OpenModalMsg tea.Model

// CloseModalMsg ferme la fenêtre modale
type CloseModalMsg struct{}

type SecurityIncreaseMsg struct{}

type SecurityScanMsg struct{}

type SecurityKickMsg struct{}

func (c *Client) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		c.height = msg.Height
		c.width = msg.Width
		return c, nil

	case OpenModalMsg:
		c.modal = msg
		c.input.Focus = false
		return c, nil

	case CloseModalMsg:
		c.modal = nil
		c.input.Focus = true
		return c, nil

	case ResultMsg:
		if msg.Error != nil {
			c.output = msg.Error.Error()
			return c, nil
		}
		c.output = msg.Output
		return c, nil

	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			// quitter l'application client
			return c, tea.Quit
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

func (c *Client) View() string {
	// affichage par défaut
	if c.modal == nil {
		return lg.JoinVertical(lg.Left,
			c.statusView(),
			c.outputView(),
			c.inputView(),
		)
	}

	// affichage dans le cas ou une fenêtre modale est ouverte
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

// Run parse et exécute la commande saisie par l'utilisateur
func (c *Client) Run() tea.Cmd {
	args := strings.Fields(c.input.Value)

	// construire la tea.Cmd qui parse et exécute la commande
	return func() tea.Msg {
		// exécuter la commande
		return c.Console.Run(c, args)
	}
}

func (c Client) Security(t time.Time) tea.Msg {
	fmt.Println("sec")
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

func (c Client) StartSecurity() tea.Cmd {
	if c.Console.Alarm == 1 {
		return tea.Every(time.Second, c.Security)
	}
	return nil
}
