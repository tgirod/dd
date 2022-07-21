package main

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
	"text/tabwriter"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

type Client struct {
	width   int    // largeur de l'affichage
	height  int    // hauteur de l'affichage
	input   Input  // invite de commande
	output  string // résultat de la dernière commande
	lastCmd string // dernière commande saisie

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
	return tea.Every(time.Second, func(t time.Time) tea.Msg {
		return SecurityMsg{}
	})
}

// affiche le résultat d'une commande
type ResultMsg struct {
	Error  error
	Output string
}

type SecurityMsg struct{}

func (c *Client) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		c.height = msg.Height
		c.width = msg.Width
		return c, nil

	case ResultMsg:
		b := strings.Builder{}
		if msg.Error != nil {
			fmt.Fprintf(&b, "%s\n\n", msg.Error.Error())
		}
		fmt.Fprintf(&b, "%s\n", msg.Output)
		c.output = b.String()
		return c, nil

	case SecurityMsg:
		return c, tea.Every(c.SecurityDelay(), c.Security)

	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			// quitter l'application client
			return c, tea.Quit
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

// temps entre deux scans de sécurité - plus long si le hacker a activé la DNI
func (c *Client) SecurityDelay() time.Duration {
	delay := time.Second
	if c.Console.DNI {
		delay *= 3
	}
	return delay
}

func (c *Client) View() string {
	return lg.JoinVertical(lg.Left,
		c.statusView(),
		// c.debugView(),
		c.outputView(),
		c.inputView(),
	)
}

var (
	// barre d'état
	statusStyle = lg.NewStyle().
			Padding(0, 1, 0, 1).
			Margin(0, 1, 0, 1).
			Foreground(lg.Color("0")).
			Background(lg.Color("10"))
	// historique
	histStyle = lg.NewStyle().
			Padding(0, 1, 0, 1).
			Margin(0, 1, 0, 1).
			Foreground(lg.Color("10")).
			Background(lg.Color("0"))

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
	login := fmt.Sprintf("👤[%s]", c.Console.Login)
	priv := fmt.Sprintf("✪[%d]", c.Console.Privilege)
	alarm := fmt.Sprintf("💀[%d]", c.Console.Alert)

	b := strings.Builder{}
	if len(c.Console.History) == 0 {
		b.WriteString("déconnecté")
	}
	for _, h := range c.Console.History {
		fmt.Fprintf(&b, "%s/", h.Address)
	}
	hist := fmt.Sprintf("🖧[%s]", b.String())

	left := fmt.Sprintf("%s %s %s ", login, priv, alarm)
	histWidth := c.width - statusStyle.GetHorizontalFrameSize() - lg.Width(left)
	status := left + lg.PlaceHorizontal(histWidth, lg.Left, hist)
	return statusStyle.Render(status)
}

func (c Client) debugView() string {
	//Alain : debug Stack
	hist := c.Console.History.AsString()

	width := c.width - statusStyle.GetHorizontalFrameSize()
	height := 5

	content := hist
	// wrap au cas ou certaines lignes seraient trop longues
	content = wordwrap.String(content, width)
	// disposer le texte dans un espace qui remplit l'écran
	content = lg.Place(width, height, lg.Left, lg.Top, content)

	return histStyle.Render(content)
}

func (c Client) outputView() string {
	// dimensions de l'espace d'affichage
	width := c.width - outputStyle.GetHorizontalFrameSize()
	// Need vertical space for debug
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
	// tenter d'augementer le niveau d'alarme
	if c.Console.Alert > 0 && rand.Float64() < c.Console.Server.Detection {
		c.Console.Alert++
	}

	if c.Console.Alert > 5 {
		// hacker repéré, il se fait kicker du serveur
		c.Console.Server = nil
		c.Console.Login = ""
		c.Console.Privilege = 0
		c.Console.Alert = 0
		c.Console.History.Clear()

		return ResultMsg{
			Output: "coupure forcée de la connexion",
		}
	}

	// on continue de faire tourner la routine de sécurité
	return SecurityMsg{}
}

func tw(output io.Writer) *tabwriter.Writer {
	return tabwriter.NewWriter(output, 8, 1, 2, ' ', 0)
}
