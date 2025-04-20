package bubble

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const numSamples int = 60 * 2
const pollTime time.Duration = time.Millisecond * 500

type Message struct {
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		}
	case Message:
		if len(m.last) > numSamples-1 {
			m.last = m.last[1:]
		}

		m.last = append(m.last, m.bat.PowerNow)
		return m, tea.Tick(pollTime, m.Monitor)
	}

	return m, tea.Tick(pollTime, m.Monitor)
}

func (m Model) View() string {
	return m.bat.String() +
		"\t" +
		m.LastStats()
}

func RunBubble() {
	p := tea.NewProgram(NewModel())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
