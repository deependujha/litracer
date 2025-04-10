package bubbletea

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	padding  = 2
	maxWidth = 80
)

type itemMsg int
type workDone int

func StartProgressBar(totalItems int, itemChan chan int) bool {
	m := progressModel{
		progress:       progress.New(progress.WithDefaultGradient()),
		itemsProcessed: 0,
		totalItems:     totalItems,
	}

	p := tea.NewProgram(m)
	go func() {
		for val := range itemChan {
			p.Send(itemMsg(val))
		}
		time.Sleep(1 * time.Second) // wait for a second before quitting for progress bar to finish
		p.Send(tea.Quit()) // gracefully shut down after duration
	}()

	if _, err := p.Run(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}

	return m.stopped
}

type progressModel struct {
	progress       progress.Model
	itemsProcessed int
	totalItems     int
	stopped        bool
}

func (m progressModel) Init() tea.Cmd {
	return nil
}

func (m progressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.stopped = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case tea.WindowSizeMsg:
		m.progress.Width = msg.Width - padding*2 - 6
		if m.progress.Width > maxWidth {
			m.progress.Width = maxWidth
		}
		return m, nil

	case workDone:
		fmt.Println("Done")
		return m, tea.Quit

	case itemMsg:
		m.itemsProcessed++
		percentage_of_items_processed := float64(m.itemsProcessed)/float64(m.totalItems) + 1e-6
		percentage_of_items_processed = min(percentage_of_items_processed, 1.0)
		// fmt.Println("percentage: ", percentage_of_items_processed, m.itemsProcessed, totalItems)
		cmd := m.progress.SetPercent(percentage_of_items_processed)
		return m, cmd

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd

	default:
		return m, nil
	}
}

func (m progressModel) View() string {
	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + m.progress.View() + fmt.Sprintf(" %d/%d\n\n", m.itemsProcessed, m.totalItems)
}
