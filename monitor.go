package main

import (
	"fmt"
	//"io"
	//"strings"
	//"text/tabwriter"
	"time"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	lg "github.com/charmbracelet/lipgloss"
	//"github.com/muesli/reflow/wordwrap"

)
const sizeConnexionMonitor = 5
const sizeRegisterMonitor = 4

// to list the register to monitor
type DescRegisterMonitor struct {
	server string
	id     int
}
var regMonitored = []DescRegisterMonitor {
	{"dd.local", 1},
}

type Monitor struct {
	fakeUser User
	Client
	startTime time.Time

	connexions map[ssh.Session]*Console // connexions actives
}

func NewMonitor(startT time.Time,
	width, height int,
	sessions map[ssh.Session]*Console) *Monitor {

	m := &Monitor{
		fakeUser: User{
			Login: "fake",
			Server: "",
			Backdoor: false,
			Groups: []string{""},
		},
		Client:    *NewClient(width, height, true),
		startTime: startT,
		// width:     width,
		// height:    height,
		// input: Input{
		// 	Focus:       true,
		// 	Placeholder		: "help",
		// },
		// output:     viewport.New(width, height-2-sizeConnexionMonitor),
		// Game:       game,
		// Console:    NewConsole(),
		connexions: sessions,
	}

	m.Client.output = viewport.New(width,
		height-2-sizeConnexionMonitor-sizeRegisterMonitor)
	m.Client.output.Style = outputStyle
	return m
}

// Pour avoir une mise à jour toutes les secondes, on utilise tea.every
// Cette fonction envoie un Message toute les secondes, mais il faut
// la relancer à chaque fois.
type TickMsg time.Time

func tickEvery() tea.Cmd {
	return tea.Every(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

// Needed to implement tea.Model
func (m *Monitor) Init() tea.Cmd {
	// On commencer à lancer des Tikcs
	return tickEvery()
}

// // affiche le résultat d'une commande
// type ResultMsg struct {
// 	Error   error
// 	Cmd     string
// 	Output  string
// 	Illegal bool
// }

// type SecurityMsg struct {
// 	Wait time.Duration // temps avant de relancer la routine de sécurité
// }

// func (m *Monitor) Wrap(output string) string {
// 	w := m.output.Width - outputStyle.GetHorizontalFrameSize()
// 	return wordwrap.String(output, w)
// }

// Update est appelé quand des tea.Msg sont reçus
// après un Update, View est appelé.
func (m *Monitor) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	//var cmd tea.Cmd
	switch msg := msg.(type) {

	case TickMsg:
		// On recommence à Ticker
		return m, tickEvery()

	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		m.output.Height = msg.Height - 2 - sizeConnexionMonitor - sizeRegisterMonitor
		return m, nil

	default:
		// Pass on to Client
		_, cmdClient := m.Client.Update(msg)
		cmds = append(cmds, cmdClient)

	// case ResultMsg:
	// 	// mettre à jour la sortie
	// 	b := strings.Builder{}
	// 	if msg.Error != nil {
	// 		fmt.Fprintf(&b, "%s\n\n", errorTextStyle.Render(msg.Error.Error()))
	// 	}
	// 	fmt.Fprintf(&b, "> %s\n\n", msg.Cmd)
	// 	fmt.Fprintf(&b, "%s\n", msg.Output)
	// 	// curOutput := m.Wrap(b.String())
	// 	curOutput := b.String()
	// 	if m.prevOutput == "" {
	// 		m.output.SetContent(curOutput)
	// 	} else {
	// 		m.output.SetContent(m.prevOutput + "\n" + curOutput)
	// 	}
	// 	m.output.GotoBottom()
	// 	m.prevOutput = curOutput

	// 	// // déclencher le scan si la commande est illégale
	// 	// if msg.Illegal {
	// 	// 	return c, m.StartSecurity
	// 	// }

	// 	return m, nil

	// 	// case SecurityMsg:
	// 	// 	if m.Console.Alert {
	// 	// 		// l'alerte est toujours là, on relance la routine de sécurité pour un tour
	// 	// 		return c, tea.Every(msg.Wait, m.Security)
	// 	// 	}

	// case tea.KeyMsg:
	// 	if msg.Type == tea.KeyCtrlC {
	// 		// quitter l'application client
	// 		return m, tea.Quit
	// 	}

	// 	if msg.Type == tea.KeyEnter {
	// 		// valider la commande

	// 		cmd = m.Run()
	// 		m.input.Value = ""
	// 		return m, cmd
	// 	}

	// 	// viewport
	// 	output, cmdOutput := m.output.Update(msg)
	// 	m.output = output

	// 	// laisser le prompt gérer
	// 	input, cmdInput := m.input.Update(msg)
	// 	m.input = input.(Input)
	// 	return m, tea.Batch(cmdOutput, cmdInput
	//	)
	}

	return m, tea.Batch(cmds...)
}

func (m *Monitor) View() string {
	return lg.JoinVertical(lg.Left,
		m.connectionsView(),
		m.Client.View(),
		// m.statusView(),
		// // c.debugView(),
		// m.output.View(),
		// m.inputView(),
	)
}

// func (m Monitor) statusView() string {
// 	now := time.Now()
// 	clock_str := fmt.Sprintf(" %02d:%02d:%02d", now.Hour(), now.Minute(), now.Second())
// 	ontime := now.Sub(m.startTime)
// 	ontime_str := fmt.Sprintf(" onTime=[%02d:%02d:%02d]",
// 		int(ontime.Hours()), int(ontime.Minutes()), int(ontime.Seconds()))
// 	//conn_str := fmt.Sprintf(" %02d co", len(m.connexions))

// 	statusWidth := m.width - statusStyle.GetHorizontalFrameSize()
// 	//fmt.Println("statusWidth", statusWidth)
// 	status := lg.PlaceHorizontal(statusWidth, lg.Left, clock_str+ontime_str)

// 	return statusStyle.Inline(true).Render(status)
// }

//var xxx = lg.NewStyle()

// func (c Client) debugView() string {
// 	//Alain : debug Stack
// 	hist := c.Console.History.AsString()

// 	width := c.width - statusStyle.GetHorizontalFrameSize()
// 	height := sizeConnexionMonitor

// 	content := hist
// 	// wrap au cas ou certaines lignes seraient trop longues
// 	content = wordwrap.String(content, width)
// 	// disposer le texte dans un espace qui remplit l'écran
// 	content = lg.Place(width, height, lg.Left, lg.Top, content)

// 	return histStyle.Render(content)
// }

func (m Monitor) connectionsView() string {
	// dimensions de l'espace d'affichage
	width := m.width - outputStyle.GetHorizontalFrameSize()
	// height := m.height - outputStyle.GetVerticalFrameSize()

	// fmt.Println("CONN width", width)
	// fmt.Println("CONN heigh", height)
	content := ""
	content += invertTextStyle.Width(m.width).Render(conHeader) + "\n"
	// liste les connexions
	for _, e := range m.connexions {
		content += fmtConnexion(e) + "\n"
	}

	// As Register is build at each query, must loop each time
	content += invertTextStyle.Width(m.width).Render(regHeader) + "\n"
	for _, desc := range regMonitored {
		serv, err := FindServer( desc.server )
		if err != nil {
			app.Log( "WARN register monitor : cannot find "+desc.server )
		} else {
			var reg Register
			//reg, err = serv.Register(desc.id, m.fakeUser)
			reg, err = serv.Register(desc.id, m.Client.Console.User)
			if err != nil {
				msg := fmt.Sprintf("WARN register %d not found on %s",
					desc.id, desc.server)
				app.Log( msg )
			} else {
				content += fmtRegister(reg) + "\n"
			}
		}
	}

	// wrap au cas ou certaines lignes seraient trop longues
	//content = wordwrap.String(content, width)

	// disposer le texte dans un espace qui remplit l'écran
	content = lg.Place(width, sizeConnexionMonitor+sizeRegisterMonitor,
		lg.Left, lg.Top, content)

	return outputStyle.Render(content)
}

// func (m Monitor) inputView() string {
// 	content := m.input.View()
// 	width := m.width - inputStyle.GetHorizontalFrameSize()
// 	content = lg.PlaceHorizontal(width, lg.Left, "> "+content)
// 	return inputStyle.Render(content)
// }

// // Run parse et exécute la commande saisie par l'utilisateur
// func (m *Monitor) Run() tea.Cmd {
// 	args := strings.Fields(m.input.Value)

// 	// construire la tea.Cmd qui parse et exécute la commande
// 	return func() tea.Msg {
// 		// exécuter la commande
// 		return m.Console.Run(&Client(*m), args)
// 	}
// }

// func (c Client) Delay() time.Duration {
// 	if c.Console.DNI {
// 		return time.Second * DNISpeed
// 	} else {
// 		return time.Second
// 	}
// }

// func (c Client) StartSecurity() tea.Msg {
// 	if !c.Console.Alert {
// 		c.Console.Alert = true
// 		c.Console.Countdown = c.Console.Server.Scan
// 		return SecurityMsg{c.Delay()}
// 	} else {
// 		// l'alerte est déjà activée, l'avancer
// 		malus := (c.Scan / 10).Round(time.Second)
// 		c.Countdown = c.Countdown - malus
// 	}

// 	return nil
// }

// func (c Client) Security(t time.Time) tea.Msg {
// 	// décrémenter d'une seconde
// 	c.Countdown -= time.Second

// 	if c.Countdown > 0 {
// 		// on continue de faire tourner la routine de sécurité
// 		return SecurityMsg{c.Delay()}
// 	}

// 	c.Disconnect()
// 	if c.DNI {
// 		return ResultMsg{
// 			Output: `
// 			     DUMPSHOCK !!!!
//                      _____
//                     /     \
//                    | () () |
//                     \  ^  /
//                      |||||
//                      |||||

// 			PERDS UN POINT DE VIE

// coupure de la connexion au réseau.
// `,
// 		}
// 	}

// 	return ResultMsg{
// 		Output: `
// coupure de la connexion au réseau.
// `,
// 	}
// }

// func (m Monitor) Disconnect() {
// 	m.Console.Server = nil
// 	m.Console.Login = ""
// 	m.Console.Privilege = 0
// 	m.Console.Alert = false
// 	m.Console.History.Clear()
// 	if len(m.Console.Sub) > 11 {
// 		m.Console.Sub = m.Console.Sub[0:11]
// 	}
// 	m.Console.MakeMatrix()
// }

// func tw(output io.Writer) *tabwriter.Writer {
// 	return tabwriter.NewWriter(output, 8, 1, 2, ' ', 0)
// }
