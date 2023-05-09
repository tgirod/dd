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

func NewClient(width, height int) *Client {
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
		Console: NewConsole(),
	}
	return c
}

func (c *Client) Init() tea.Cmd {
	return tea.Batch(
		textinput.Blink,  // clignottement du curseur
		c.TickSecurity(), // routine de sécurité
	)
}

type SecurityMsg struct{}

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
	case tea.WindowSizeMsg:
		// gestion de la taille de fenêtre
		c.width = msg.Width
		c.height = msg.Height
		c.status.Width = msg.Width
		c.output.Width = msg.Width
		c.output.Height = msg.Height - 2
		c.input.Width = msg.Width

	case OpenModalMsg:
		// ouvrir une fenêtre modale
		cmd = c.OpenModal(msg.(tea.Model))
		cmds = append(cmds, cmd)

	case CloseModalMsg:
		// fermer une fenêtre modale
		cmds = append(cmds, c.CloseModal())

	case Result:
		// afficher le résultat d'une commande
		c.Console.AddResult(msg)
		c.RenderOutput()

	case Context:
		// reprendre l'exécution d'un contexte
		cmd := func() tea.Msg {
			return msg.Resume([]string{})
		}
		cmds = append(cmds, cmd)

	case SecurityMsg:
		if c.Console.Alert {
			c.Console.TickAlert()
		}
		cmds = append(cmds, c.TickSecurity())

	case tea.KeyMsg:
		// gestion du clavier
		if c.modal != nil {
			break
		}

		switch msg.Type {
		case tea.KeyCtrlC:
			// quitter l'application client
			cmds = append(cmds, tea.Quit)

		case tea.KeyEnter:
			// valider la commande
			prompt := c.input.Value()          // récupérer le prompt
			c.input.Reset()                    // effacer le champ
			msg := c.Parse(prompt)             // exécuter et récupérer le résultat
			cmds = append(cmds, MsgToCmd(msg)) // injecter le résultat dans la boucle

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
		// passer tous les messages au prompt
		if c.modal != nil {
			break
		}

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
	login := c.Console.Identity.Login

	groups := "public"
	if len(c.Groups) > 0 {
		groups = strings.Join(c.Groups, " ")
	}

	timer := "--:--"
	if c.Console.Alert {
		min := int(c.Countdown.Minutes())
		sec := int(c.Countdown.Seconds()) - min*60
		timer = fmt.Sprintf("%02d:%02d", min, sec)
	}

	// chemin de connexion
	hist := "déconnecté"
	if c.IsConnected() {
		hist = c.Console.Session.Path()
	}

	c.status.SetContent(timer, hist, login, groups)

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
	for _, e := range c.Console.Results {
		if e.Prompt != "" {
			fmt.Fprintf(&b, "> %s\n\n",
				promptStyle.MaxWidth(c.width).Render(e.Prompt))
		}

		if e.Error != nil {
			fmt.Fprintf(&b, "%s\n\n",
				errorStyle.MaxWidth(c.width).Render(e.Error.Error()))
		}

		if e.Output != "" {
			fmt.Fprintf(&b, "%s\n\n",
				outputStyle.MaxWidth(c.width).Render(e.Output))
		}
	}

	c.output.SetContent(b.String())
	c.output.GotoBottom()
}

// TickSecurity déclenche un scan de sécurité
// cette méthode est appelée toutes les secondes
func (c *Client) TickSecurity() tea.Cmd {
	return tea.Every(c.Delay(), func(t time.Time) tea.Msg {
		return SecurityMsg{}
	})
}

func (c *Client) OpenModal(model tea.Model) tea.Cmd {
	c.input.Blur()
	initCmd := model.Init()
	// initialiser la taille si nécessaire
	w, h := c.modalWindowSize()
	var sizeCmd tea.Cmd
	model, sizeCmd = model.Update(tea.WindowSizeMsg{Width: w, Height: h})
	c.modal = model
	return tea.Batch(
		initCmd,
		sizeCmd,
	)
}

func (c *Client) CloseModal() tea.Cmd {
	c.modal = nil
	return c.input.Focus()
}

func tw(output io.Writer) *tabwriter.Writer {
	return tabwriter.NewWriter(output, 0, 0, 2, ' ', 0)
}

func MsgToCmd(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}
