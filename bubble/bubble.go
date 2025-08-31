package bubble

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
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
			return m, tea.Quit

		}
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
	return m.bat.String() +
		"\t" +
		m.LastStats()
}

func RunBubble() {
	p := tea.NewProgram(NewApp())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
