package bubble

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const numSamples int = 4 * 60 * 2
const pollTime time.Duration = time.Millisecond * 250

type AppMessage struct {
}

func (m App) Init() tea.Cmd {
	return nil
}

func (m App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quit = true
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width

		return m, tea.Tick(pollTime, m.Monitor)
	case AppMessage:
		if len(m.last) > numSamples-1 {
			m.last = m.last[1:]
		}

		m.last = append(m.last, m.bat.PowerNow)
		return m, tea.Tick(pollTime, m.Monitor)
	}

	return m, tea.Tick(pollTime, m.Monitor)
}

func (m App) View() string {
	if m.quit {
		return ""
	}

	legacy := m.bat.RenderStats() + "\t" + m.LastStats()

	rendered := BorderedStyle.Render(legacy)

	width, _ := lipgloss.Size(rendered)

	if width > m.width {
		rendered = BorderedStyle.Render(m.bat.RenderStats())
	}
	return rendered
}

func RunBubble() {
	p := tea.NewProgram(NewApp(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
