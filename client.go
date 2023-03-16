package main

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
	"github.com/knipferrc/teacup/statusbar"
)

const DNISpeed = 3

type Client struct {
	width      int              // largeur de l'affichage
	height     int              // hauteur de l'affichage
	input      textinput.Model  // invite de commande
	output     viewport.Model   // affichage de la sortie des commandes
	status     statusbar.Bubble // barre de statut
	prevOutput string           // sortie de la commande précédente
	modal      tea.Model        // fenêtre modale

	*Game    // état interne du jeu
	*Console // console enregistrée dans le jeu
}

func NewClient(width, height int, game *Game) *Client {
	// barre de statut
	status := statusbar.New(
		statusbar.ColorConfig{
			Foreground: lg.AdaptiveColor{Dark: "#ffffff", Light: "#ffffff"},
			Background: lg.AdaptiveColor{Light: "#F25D94", Dark: "#F25D94"},
		},
		statusbar.ColorConfig{
			Foreground: lg.AdaptiveColor{Light: "#ffffff", Dark: "#ffffff"},
			Background: lg.AdaptiveColor{Light: "#3c3836", Dark: "#3c3836"},
		},
		statusbar.ColorConfig{
			Foreground: lg.AdaptiveColor{Light: "#ffffff", Dark: "#ffffff"},
			Background: lg.AdaptiveColor{Light: "#A550DF", Dark: "#A550DF"},
		},
		statusbar.ColorConfig{
			Foreground: lg.AdaptiveColor{Light: "#ffffff", Dark: "#ffffff"},
			Background: lg.AdaptiveColor{Light: "#6124DF", Dark: "#6124DF"},
		},
	)
	status.SetSize(width)

	// zone d'affichage des résultats
	output := viewport.New(width, height-2)

	// prompt
	input := textinput.New()
	input.Width = width
	input.Focus()

	c := &Client{
		width:   width,
		height:  height,
		input:   input,
		output:  output,
		status:  status,
		Game:    game,
		Console: NewConsole(),
	}
	return c
}

var startSecurity = tea.Every(time.Second, func(t time.Time) tea.Msg {
	return SecurityMsg{}
})

func (c *Client) Init() tea.Cmd {
	return tea.Batch(
		startSecurity,
		textinput.Blink,
	)
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

type OpenModalMsg tea.Model

type CloseModalMsg struct{}

func (c *Client) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		// redimensionner les différentes parties de l'interface
		c.width = msg.Width
		c.height = msg.Height
		c.status.Width = msg.Width
		c.output.Width = msg.Width
		c.output.Height = msg.Height - 2
		c.input.Width = msg.Width
		if c.modal != nil {
			c.modal, cmd = c.modal.Update(msg)
			cmds = append(cmds, cmd)
		}

	case OpenModalMsg:
		// ouvrir une fenêtre modale
		c.modal = msg.(tea.Model)
		cmd = c.modal.Init()
		c.input.Blur()
		cmds = append(cmds, cmd)

	case CloseModalMsg:
		c.modal = nil
		c.input.Focus()
		cmds = append(cmds, textinput.Blink)

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
			cmds = append(cmds, c.StartSecurity)
		}

	case SecurityMsg:
		if c.Console.Alert {
			// l'alerte est toujours là, on relance la routine de sécurité pour un tour
			cmds = append(cmds, tea.Every(msg.Wait, c.Security))
		}

	case tea.KeyMsg:
		if c.modal != nil {
			// la fenêtre modale prend le contrôle du clavier
			c.modal, cmd = c.modal.Update(msg)
			cmds = append(cmds, cmd)
		} else {
			switch msg.Type {
			case tea.KeyCtrlC:
				// quitter l'application client
				cmds = append(cmds, tea.Quit)
			case tea.KeyEnter:
				// valider la commande
				input := c.input.Value()
				c.input.Reset()
				cmd = c.Parse(input)
				cmds = append(cmds, cmd)
			case tea.KeyPgUp, tea.KeyPgDown:
				// scroll de la sortie
				c.output, cmd = c.output.Update(msg)
				cmds = append(cmds, cmd)
			default:
				// passer le KeyMsg au prompt
				c.input, cmd = c.input.Update(msg)
				cmds = append(cmds, cmd)
			}
		}

	default:
		if c.modal != nil {
			// la fenêtre modale prend le contrôle du clavier
			c.modal, cmd = c.modal.Update(msg)
			cmds = append(cmds, cmd)
		} else {
			// passer tous les messages au prompt
			c.input, cmd = c.input.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	return c, tea.Batch(cmds...)
}

func (c *Client) View() string {
	c.statusView() // mettre à jour la barre de statut

	if c.modal != nil {
		return c.modal.View()
	}

	return lg.JoinVertical(lg.Left,
		c.status.View(),
		c.output.View(),
		c.input.View(),
	)
}

var (
	// fenêtre modale
	modalStyle = lg.NewStyle().
			Padding(0, 1, 0, 1).
			Margin(0, 1, 0, 1).
			BorderStyle(lg.DoubleBorder()).
			BorderForeground(lg.Color("10"))

	// texte erreur
	errorTextStyle = lg.NewStyle().Foreground(lg.Color("9")).Padding(1)
)

func (c *Client) statusView() {
	login := fmt.Sprintf("id=%s", c.Console.Login)
	priv := fmt.Sprintf("priv=%d", c.Console.Privilege)
	timer := "--:--"
	if c.Console.Alert {
		min := int(c.Countdown.Minutes())
		sec := int(c.Countdown.Seconds()) - min*60
		timer = fmt.Sprintf("%02d:%02d", min, sec)
	}

	// historique complet
	b := strings.Builder{}
	if len(c.Console.History) == 0 {
		b.WriteString("déconnecté")
	}
	for _, h := range c.Console.History {
		fmt.Fprintf(&b, "%s@%s/", h.Login, h.Address)
	}

	hist := fmt.Sprintf("net=%s", b.String())

	c.status.SetContent(
		timer,
		hist,
		login,
		priv,
	)
}

// Run parse et exécute la commande saisie par l'utilisateur
func (c *Client) Parse(input string) tea.Cmd {
	args := strings.Fields(input)

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
