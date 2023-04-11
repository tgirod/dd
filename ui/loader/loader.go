package loader

import (
	"math"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type tickMsg time.Time

type LoadedMsg struct{}

type Model struct {
	msg      tea.Msg        // message a envoyer une fois le chargement terminé
	progress progress.Model // barre de chargement
	duration time.Duration
	status   []string
}

func New(msg tea.Msg, duration time.Duration, status []string) *Model {
	m := Model{
		msg:      msg,
		progress: progress.New(progress.WithDefaultGradient()),
		duration: duration,
		status:   status,
	}
	m.progress.SetSpringOptions(1, 0)
	return &m
}

func (m *Model) tickCmd() tea.Cmd {
	return tea.Every(time.Second,
		func(t time.Time) tea.Msg {
			return tickMsg(t)
		},
	)
}

func (m *Model) Init() tea.Cmd {
	return m.tickCmd()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width

	case tickMsg:
		if m.progress.Percent() >= 1.0 {
			cmds = append(cmds,
				func() tea.Msg { return m.msg },
				func() tea.Msg { return LoadedMsg{} },
			)
			break
		}

		inc := 1.0 / m.duration.Seconds()
		cmds = append(cmds,
			m.progress.IncrPercent(inc),
			m.tickCmd(),
		)

	default:
		prog, cmd := m.progress.Update(msg)
		m.progress = prog.(progress.Model)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) View() string {
	status := "exécution en cours ..."
	if len(m.status) > 0 {
		index := int(math.Floor(float64(len(m.status)) * m.progress.Percent()))
		if index == len(m.status) {
			index--
		}
		status = m.status[index]
	}

	return lg.JoinVertical(lg.Center,
		m.progress.View(),
		status,
	)
}
