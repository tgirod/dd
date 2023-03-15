package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
	"github.com/knipferrc/teacup/statusbar"
)

type Client struct {
	width, height int              // dimensions de la fenêtre
	status        statusbar.Bubble // barre de statut en haut de la fenêtre
	output        viewport.Model   // zone d'affichage des résultats
	prompt        textinput.Model  // zone de saisie de commande
	modal         tea.Model        // fenêtre modale
}

type OpenModalMsg tea.Model

type CloseModalMsg struct{}

// DisplayMsg fournit du contenu à afficher dans le display
type DisplayMsg string

func (c *Client) Init() tea.Cmd {
	return textinput.Blink
}

func (c *Client) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	log.Printf("%T %+v\n", msg, msg)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		log.Println("client", msg)
		// dimensions de la fenêtre
		c.width = msg.Width
		c.height = msg.Height

		// dimensions de la barre de statut (H=1)
		c.status.SetSize(c.width)

		// dimensions du prompt (H=1)
		c.prompt.Width = c.width

		// dimensions de la sortie
		c.output.Width = c.width
		c.output.Height = c.height - 2

		// transmettre le message à la fenêtre modale
		if c.modal != nil {
			msg := tea.WindowSizeMsg{c.output.Width, c.output.Height}
			c.modal, cmd = c.modal.Update(msg)
			cmds = append(cmds, cmd)
		}

	case DisplayMsg:
		// affiche quelque chose dans le display
		c.output.SetContent(string(msg))
		c.output.GotoBottom()

	case OpenModalMsg:
		c.modal = msg.(tea.Model)
		cmd = c.modal.Init()
		c.prompt.Blur()
		cmds = append(cmds, cmd)

	case CloseModalMsg:
		c.modal = nil
		c.prompt.Focus()
		cmds = append(cmds, textinput.Blink)

	case tea.KeyMsg:
		if c.modal != nil {
			// la fenêtre modale gère le clavier
			c.modal, cmd = c.modal.Update(msg)
			cmds = append(cmds, cmd)
		} else {
			switch msg.Type {
			case tea.KeyCtrlC:
				// quitter l'application de force
				cmds = append(cmds, tea.Quit)
			case tea.KeyPgUp, tea.KeyPgDown:
				// scroll l'affichage
				c.output, cmd = c.output.Update(msg)
				cmds = append(cmds, cmd)
			case tea.KeyEnter:
				// valider une commande
				input := c.prompt.Value()
				c.prompt.Reset()
				// retourner une Cmd pour lancer le parsing
				cmds = append(cmds, c.Parse(input))
			default:
				// envoyer vers le prompt
				c.prompt, cmd = c.prompt.Update(msg)
				cmds = append(cmds, cmd)
			}
		}

	default:
		if c.modal != nil {
			c.modal, cmd = c.modal.Update(msg)
			cmds = append(cmds, cmd)
		}

		c.prompt, cmd = c.prompt.Update(msg)
		cmds = append(cmds, cmd)
	}

	return c, tea.Batch(cmds...)
}

func (c *Client) Parse(input string) tea.Cmd {
	return func() tea.Msg {
		if input == "mod" {
			mod := NewModal(c.output.Width, c.output.Height)
			log.Println("client", "NewModal", mod)
			return OpenModalMsg(mod)
		}
		return DisplayMsg(input)
	}
}

func (c *Client) View() string {
	if c.modal != nil {
		return lg.JoinVertical(lg.Left,
			c.status.View(),
			c.modal.View(),
			c.prompt.View(),
		)
	}

	return lg.JoinVertical(lg.Left,
		c.status.View(),
		c.output.View(),
		c.prompt.View(),
	)
}

func NewClient() *Client {
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
	status.SetContent("A", "B", "C", "D")
	output := viewport.New(0, 0)
	prompt := textinput.New()
	prompt.Focus()

	c := Client{
		status: status,
		output: output,
		prompt: prompt,
	}
	return &c
}

type Modal struct {
	width, height int
	keymap        ModalKeymap
}

type ModalKeymap struct {
	Quit key.Binding
}

func NewModal(width, height int) *Modal {
	mod := new(Modal)
	mod.width = width
	mod.height = height
	mod.keymap = ModalKeymap{
		Quit: key.NewBinding(
			key.WithKeys("q"), key.WithHelp("q", "quitter"),
		),
	}
	return mod
}

func (m *Modal) Init() tea.Cmd {
	return nil
}

func (m *Modal) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		log.Println("modal", msg)
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.Quit):
			return m, func() tea.Msg { return CloseModalMsg{} }
		}
	}

	return m, nil
}

var modalStyle = lg.NewStyle().
	Border(lg.NormalBorder()).
	Padding(1).
	Margin(4)

func (m *Modal) View() string {
	str := "fenêtre modale : q pour quitter"
	return lg.Place(m.width, m.height, lg.Center, lg.Center, modalStyle.Render(str))
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		fmt.Println("fatal:", err)
		os.Exit(1)
	}
	defer f.Close()

	p := tea.NewProgram(NewClient(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}