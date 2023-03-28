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
		Console: NewConsole(game),
	}
	return c
}

func (c *Client) Init() tea.Cmd {
	return tea.Batch(
		textinput.Blink,
	)
}

type SecurityMsg struct {
	Wait time.Duration // temps avant de relancer la routine de sécurité
}

type OpenModalMsg tea.Model

type CloseModalMsg struct{}

func (c *Client) modalWindowSize() (int, int) {
	w, h := modalStyle.GetFrameSize()
	return c.width - w, c.height - h - 1
}

func (c *Client) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case BackMsg:
		c.Console.Back()
		c.RenderOutput()

	case ConnectMsg:
		c.Console.Connect(msg.Address)
		c.RenderOutput()

	case DataSearchMsg:
		c.Console.DataSearch(msg.Keyword)
		c.RenderOutput()

	case DataViewMsg:
		c.Console.DataView(msg.Id)
		c.RenderOutput()

	case HelpMsg:
		c.Console.Help(msg.Args)
		c.RenderOutput()

	case IdentifyMsg:
		c.Console.Identify(msg.Login, msg.Password)
		c.RenderOutput()

	case IndexMsg:
		c.Console.Index()
		c.RenderOutput()

	case LinkListMsg:
		c.Console.LinkList()
		c.RenderOutput()

	case LinkMsg:
		c.Console.Link(msg.Id)
		c.RenderOutput()

	case LoadMsg:
		c.Console.Load(msg.Code)
		c.RenderOutput()

	case PlugMsg:
		c.Console.Plug()
		c.RenderOutput()

	case QuitMsg:
		c.Console.Quit()
		c.RenderOutput()

	case RegistrySearchMsg:
		c.Console.RegistrySearch(msg.Name)
		c.RenderOutput()

	case RegistryEditMsg:
		c.Console.RegistryEdit(msg.Name)
		c.RenderOutput()

	case JackMsg:
		c.Console.Jack(msg.Id)
		cmds = append(cmds, c.StartSecurity)
		c.RenderOutput()

	case EvadeListMsg:
		c.Console.EvadeList()
		cmds = append(cmds, c.StartSecurity)
		c.RenderOutput()

	case EvadeMsg:
		c.Console.Evade(msg.Zone)
		cmds = append(cmds, c.StartSecurity)
		c.RenderOutput()

	case tea.WindowSizeMsg:
		// redimensionner les différentes parties de l'interface
		c.width = msg.Width
		c.height = msg.Height
		c.status.Width = msg.Width
		c.output.Width = msg.Width
		c.output.Height = msg.Height - 2
		c.input.Width = msg.Width

	case OpenModalMsg:
		c.input.Blur()
		// ouvrir une fenêtre modale
		c.modal = msg.(tea.Model)
		cmd = c.modal.Init()
		cmds = append(cmds, cmd)
		// envoyer un WindowSizeMsg
		w, h := c.modalWindowSize()
		c.modal, cmd = c.modal.Update(tea.WindowSizeMsg{Width: w, Height: h})
		cmds = append(cmds, cmd)

	case CloseModalMsg:
		c.modal = nil
		c.input.Focus()
		cmds = append(cmds, textinput.Blink)

	case Eval:
		c.Console.AppendOutput(msg)
		c.RenderOutput()

	case SecurityMsg:
		if c.Console.Alert {
			// l'alerte est toujours là
			// la routine de sécurité continue
			cmds = append(cmds, tea.Every(msg.Wait, c.Security))
		}

	case tea.KeyMsg:
		if c.modal != nil {
			break
		}

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

	default:
		if c.modal != nil {
			break
		}

		// passer tous les messages au prompt
		c.input, cmd = c.input.Update(msg)
		cmds = append(cmds, cmd)
	}

	if c.modal != nil {
		switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			w, h := c.modalWindowSize()
			c.modal, cmd = c.modal.Update(tea.WindowSizeMsg{Width: w, Height: h})
			cmds = append(cmds, cmd)
		default:
			c.modal, cmd = c.modal.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	return c, tea.Batch(cmds...)
}

var (
	modalStyle  = lg.NewStyle().Border(lg.DoubleBorder()).Padding(1)
	errorStyle  = lg.NewStyle().Foreground(lg.Color("9"))
	promptStyle = lg.NewStyle().Foreground(lg.Color("8"))
	outputStyle = lg.NewStyle()
)

func (c *Client) View() string {
	// mise à jour de la barre de statut
	login := fmt.Sprintf("id=%s", c.Console.Login)
	admin := "user"
	if c.Console.Admin {
		admin = "admin"
	}
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
		fmt.Fprintf(&b, "%s/", h.Address)
	}

	hist := fmt.Sprintf("net=%s", b.String())

	c.status.SetContent(timer, hist, login, admin)

	if c.modal != nil {
		content := modalStyle.Render(c.modal.View())
		modal := lg.Place(c.width, c.height-1, lg.Center, lg.Center, content, lg.WithWhitespaceChars(". "))
		return lg.JoinVertical(lg.Left,
			c.status.View(),
			modal,
		)
	}

	return lg.JoinVertical(lg.Left,
		c.status.View(),
		c.output.View(),
		c.input.View(),
	)
}

func (c *Client) RenderOutput() {
	b := strings.Builder{}
	for _, e := range c.Console.Evals {
		if e.Cmd != "" {
			fmt.Fprintf(&b, "> %s\n",
				promptStyle.MaxWidth(c.width).Render(e.Cmd))
		}

		if e.Error != nil {
			fmt.Fprintf(&b, "%s\n",
				errorStyle.MaxWidth(c.width).Render(e.Error.Error()))
		}

		if e.Output != "" {
			fmt.Fprintf(&b, "%s\n",
				outputStyle.MaxWidth(c.width).Render(e.Output))
		}
	}

	c.output.SetContent(b.String())
	c.output.GotoBottom()
}

// Run parse et exécute la commande saisie par l'utilisateur
func (c *Client) Parse(input string) tea.Cmd {
	args := strings.Fields(input)

	// construire la tea.Cmd qui parse et exécute la commande
	return func() tea.Msg {
		// exécuter la commande
		return c.Console.Run(args)
	}
}

func (c *Client) Delay() time.Duration {
	if c.Console.DNI {
		return time.Second * DNISpeed
	} else {
		return time.Second
	}
}

func (c *Client) StartSecurity() tea.Msg {
	c.StartAlert()
	return SecurityMsg{c.Delay()}
}

func (c *Client) Security(t time.Time) tea.Msg {
	// décrémenter d'une seconde
	c.Countdown -= time.Second

	if c.Countdown > 0 {
		// tant que l'horloge n'est pas à arrivée à 0, on ne fait rien
		return SecurityMsg{c.Delay()}
	}

	c.Console.Disconnect()
	c.RenderOutput()
	return nil
}

func tw(output io.Writer) *tabwriter.Writer {
	return tabwriter.NewWriter(output, 8, 1, 2, ' ', 0)
}
