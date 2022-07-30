package main

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

const DNISpeed = 3

type Client struct {
	width      int            // largeur de l'affichage
	height     int            // hauteur de l'affichage
	input      Input          // invite de commande
	output     viewport.Model // affichage de la sortie des commandes
	prevOutput string         // sortie de la commande précédente

	*Game    // état interne du jeu
	*Console // console enregistrée dans le jeu
}

func NewClient(width, height int, game *Game) *Client {
	c := &Client{
		width:  width,
		height: height,
		input: Input{
			Focus:       true,
			Placeholder: "help",
		},
		output:  viewport.New(width, height-2),
		Game:    game,
		Console: NewConsole(),
	}
	c.output.Style = outputStyle
	return c
}

func (c *Client) Init() tea.Cmd {
	return tea.Every(time.Second, func(t time.Time) tea.Msg {
		return SecurityMsg{}
	})
}

// affiche le résultat d'une commande
type ResultMsg struct {
	Error   error
	Cmd     string
	Output  string
	Illegal bool
}

type SecurityMsg struct {
	Wait time.Duration // temps avant de relancer la routine de sécurité
}

func (c *Client) Wrap(output string) string {
	w := c.output.Width - outputStyle.GetHorizontalFrameSize()
	return wordwrap.String(output, w)
}

func (c *Client) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		c.height = msg.Height
		c.width = msg.Width
		c.output.Height = msg.Height - 2
		return c, nil

	case ResultMsg:
		// mettre à jour la sortie
		b := strings.Builder{}
		if msg.Error != nil {
			fmt.Fprintf(&b, "%s\n\n", errorTextStyle.Render(msg.Error.Error()))
		}
		fmt.Fprintf(&b, "> %s\n\n", msg.Cmd)
		fmt.Fprintf(&b, "%s\n", msg.Output)
		// curOutput := c.Wrap(b.String())
		curOutput := b.String()
		if c.prevOutput == "" {
			c.output.SetContent(curOutput)
		} else {
			c.output.SetContent(c.prevOutput + "\n" + curOutput)
		}
		c.output.GotoBottom()
		c.prevOutput = curOutput

		// déclencher le scan si la commande est illégale
		if msg.Illegal {
			return c, c.StartSecurity
		}

		return c, nil

	case SecurityMsg:
		if c.Console.Alert {
			// l'alerte est toujours là, on relance la routine de sécurité pour un tour
			return c, tea.Every(msg.Wait, c.Security)
		}

	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			// quitter l'application client
			return c, tea.Quit
		}

		if msg.Type == tea.KeyEnter {
			// valider la commande
			cmd = c.Run()
			c.input.Value = ""
			return c, cmd
		}

		// viewport
		output, cmdOutput := c.output.Update(msg)
		c.output = output

		// laisser le prompt gérer
		input, cmdInput := c.input.Update(msg)
		c.input = input.(Input)
		return c, tea.Batch(cmdOutput, cmdInput)
	}

	return c, cmd
}

func (c *Client) View() string {
	return lg.JoinVertical(lg.Left,
		c.statusView(),
		// c.debugView(),
		c.output.View(),
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
	timer := "😀[--:--]"
	if c.Console.Alert {
		min := int(c.Countdown.Minutes())
		sec := int(c.Countdown.Seconds()) - min*60
		timer = fmt.Sprintf("💀[%02d:%02d]", min, sec)
	}

	left := fmt.Sprintf("%s %s %s ", login, priv, timer)

	// longueur max pour l'historique
	max := c.width - statusStyle.GetHorizontalFrameSize() - lg.Width(left)

	// historique complet
	b := strings.Builder{}
	if len(c.Console.History) == 0 {
		b.WriteString("déconnecté")
	}
	for _, h := range c.Console.History {
		fmt.Fprintf(&b, "%s@%s/", h.Login, h.Address)
	}
	hist := []rune(fmt.Sprintf("🖧[%s]", b.String()))

	if len(hist) > max {
		hist = hist[len(hist)-max : len(hist)]
	}

	status := left + lg.PlaceHorizontal(max, lg.Left, string(hist))
	return statusStyle.Inline(true).Render(status)
}

var xxx = lg.NewStyle()

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

// func (c Client) outputView() string {
// 	// dimensions de l'espace d'affichage
// 	width := c.width - outputStyle.GetHorizontalFrameSize()
// 	// Need vertical space for debug
// 	height := c.height - 2 - outputStyle.GetVerticalFrameSize()

// 	// dernière commande + output
// 	content := ""
// 	if c.lastCmd != "" {
// 		content = lg.JoinVertical(lg.Left,
// 			fmt.Sprintf("> %s\n", c.lastCmd),
// 			c.output.View()
// 		)
// 	} else {
// 		content = c.output
// 	}

// 	// wrap au cas ou certaines lignes seraient trop longues
// 	content = wordwrap.String(content, width)

// 	// disposer le texte dans un espace qui remplit l'écran
// 	content = lg.Place(width, height, lg.Left, lg.Bottom, content)

// 	return outputStyle.Render(content)
// }

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

func (c Client) Delay() time.Duration {
	if c.Console.DNI {
		return time.Second * DNISpeed
	} else {
		return time.Second
	}
}

func (c Client) StartSecurity() tea.Msg {
	if !c.Console.Alert {
		c.Console.Alert = true
		c.Console.Countdown = c.Console.Server.Scan
		return SecurityMsg{c.Delay()}
	} else {
		// l'alerte est déjà activée, l'avancer
		malus := (c.Scan / 10).Round(time.Second)
		c.Countdown = c.Countdown - malus
	}

	return nil
}

func (c Client) Security(t time.Time) tea.Msg {
	// décrémenter d'une seconde
	c.Countdown -= time.Second

	if c.Countdown > 0 {
		// on continue de faire tourner la routine de sécurité
		return SecurityMsg{c.Delay()}
	}

	c.Disconnect()
	if c.DNI {
		return ResultMsg{
			Output: `
			     DUMPSHOCK !!!!
                     _____
                    /     \
                   | () () |
                    \  ^  /
                     |||||
                     |||||

			PERDS UN POINT DE VIE

coupure de la connexion au réseau.
`,
		}
	}

	return ResultMsg{
		Output: `
coupure de la connexion au réseau.
`,
	}
}

func (c Client) Disconnect() {
	c.Console.Server = nil
	c.Console.Login = ""
	c.Console.Privilege = 0
	c.Console.Alert = false
	c.Console.History.Clear()
	if len(c.Console.Sub) > 11 {
		c.Console.Sub = c.Console.Sub[0:11]
	}
}

func tw(output io.Writer) *tabwriter.Writer {
	return tabwriter.NewWriter(output, 8, 1, 2, ' ', 0)
}
