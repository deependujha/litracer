package bubbletea

import (
	"fmt"
	"math/rand/v2"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	// Available spinners
	spinners = []spinner.Spinner{
		spinner.Line,
		spinner.Dot,
		spinner.MiniDot,
		spinner.Jump,
		spinner.Pulse,
		spinner.Points,
		spinner.Globe,
		spinner.Moon,
		spinner.Monkey,
	}

	textStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("252")).Render
	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	// helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render
)

func KeepSpinning(durationInSeconds int, task string) bool {
	m := model{task: task}
	m.resetSpinner()

	p := tea.NewProgram(&m)
	go func() {
		time.Sleep(time.Duration(durationInSeconds) * time.Millisecond)
		p.Send(tea.Quit()) // gracefully shut down after duration
	}()

	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}

	return m.stopped // return true, if you can run spinner again
}

type model struct {
	index   int
	spinner spinner.Model
	stopped bool
	task    string
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.stopped = true
			return m, tea.Quit
		case "h", "left":
			m.index--
			if m.index < 0 {
				m.index = len(spinners) - 1
			}
			m.resetSpinner()
			return m, m.spinner.Tick
		case "l", "right":
			m.index++
			if m.index >= len(spinners) {
				m.index = 0
			}
			m.resetSpinner()
			return m, m.spinner.Tick
		default:
			return m, nil
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	default:
		return m, nil
	}
}

func (m *model) resetSpinner() {
	random_index := rand.IntN(len(spinners)) // any random index
	m.index = random_index
	m.spinner = spinner.New()
	m.spinner.Style = spinnerStyle
	m.spinner.Spinner = spinners[m.index]
}

func (m model) View() (s string) {
	var gap string
	switch m.index {
	case 1:
		gap = ""
	default:
		gap = " "
	}

	s += fmt.Sprintf("%s%s%s", m.spinner.View(), gap, textStyle(m.task, "..."))
	// s += helpStyle("h/l, ←/→: change spinner • q: exit\n")
	return
}
