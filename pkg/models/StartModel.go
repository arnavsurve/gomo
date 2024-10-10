package models

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/stopwatch"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	errMsg  error
	tickMsg time.Time
)

type startModel struct {
	stopwatch stopwatch.Model
	spinner   spinner.Model
	progress  progress.Model
	duration  int
	quitting  bool
	err       error
	conf      Config
}

// NewStartModel displays the pomodoro timer's progress. It takes a duration in seconds.
func NewStartModel(duration int) startModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	p := progress.New(progress.WithDefaultGradient())

	sw := stopwatch.New()

	c := Config{}

	return startModel{spinner: s, progress: p, stopwatch: sw, duration: duration, conf: *c.GetConf()}
}

func (m startModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		m.stopwatch.Init(),
		tickCmd(),
	)
}

func (m startModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case tickMsg:
		if m.progress.Percent() == 1.0 {
			return m, tea.Quit // TODO handle progress complete - sound and prompt for break?
		}

		// Note that you can also use progress.Model.SetPercent to set the
		// percentage value explicitly, too.
		// TODO scale this with pomodoro time (20min/10min/15min) (duration)
		// Increments once every second
		cmd := m.progress.IncrPercent(1 / float64(m.duration))
		return m, tea.Batch(tickCmd(), cmd)

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	case errMsg:
		m.err = msg
		return m, nil

	default:
		var cmds []tea.Cmd
		var cmd tea.Cmd

		m.spinner, cmd = m.spinner.Update(msg)
		cmds = append(cmds, cmd)

		m.stopwatch, cmd = m.stopwatch.Update(msg)
		cmds = append(cmds, cmd)

		return m, tea.Batch(cmds...)
	}
}

func (m startModel) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	str := fmt.Sprintf("\n  %s\n\n  %s %s\n\n  Ctrl-C or q to quit\n\n", m.stopwatch.Elapsed(), m.spinner.View(), m.progress.View())
	if m.quitting {
		return str + "\n"
	}
	return str
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second*1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
