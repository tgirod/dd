package main

import (
	"math"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	lg "github.com/charmbracelet/lipgloss"
)

type LoaderModel struct {
	ctx      Context
	bar      progress.Model
	duration time.Duration
	status   []string
}

func NewLoader(ctx Context, duration time.Duration, status []string) *LoaderModel {
	m := LoaderModel{
		ctx:      ctx,
		bar:      progress.New(progress.WithDefaultGradient()),
		duration: duration,
		status:   status,
	}

	m.bar.SetSpringOptions(1, 0)
	return &m
}

type tickMsg time.Time

func (m *LoaderModel) tickCmd() tea.Cmd {
	return tea.Every(time.Second,
		func(t time.Time) tea.Msg {
			return tickMsg(t)
		},
	)
}

func (m *LoaderModel) Init() tea.Cmd {
	return m.tickCmd()
}

func (m *LoaderModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.bar.Width = msg.Width

	case tickMsg:
		if m.bar.Percent() >= 1.0 {
			cmds = append(cmds,
				func() tea.Msg { return m.ctx },
				func() tea.Msg { return CloseModalMsg{} },
			)
			break
		}

		inc := 1.0 / m.duration.Seconds()
		cmds = append(cmds,
			m.bar.IncrPercent(inc),
			m.tickCmd(),
		)

	default:
		bar, cmd := m.bar.Update(msg)
		m.bar = bar.(progress.Model)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m *LoaderModel) View() string {
	status := "exécution en cours ..."
	if len(m.status) > 0 {
		index := int(math.Floor(float64(len(m.status)) * m.bar.Percent()))
		if index == len(m.status) {
			index--
		}
		status = m.status[index]
	}

	return lg.JoinVertical(lg.Center,
		m.bar.View(),
		status,
	)
}
