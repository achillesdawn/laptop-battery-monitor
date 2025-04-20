package bubble

import (
	"time"

	"github.com/achillesdawn/battery-monitor/stats"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	bat *stats.BatStats
}

type Message struct {
	content string
}

func NewModel() Model {
	bat, err := stats.New()
	if err != nil {
		panic(err)
	}

	bat.CalcTimeLeft()

	return Model{
		bat: bat,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Monitor(t time.Time) tea.Msg {
	m.bat.ReadPowerAndEnergy()

	m.bat.CalcTimeLeft()

	return Message{
		content: m.bat.String(),
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		}
	case Message:
		return m, tea.Tick(time.Millisecond*500, m.Monitor)
	}

	return m, tea.Tick(time.Millisecond*500, m.Monitor)
}

func (m Model) View() string {
	return m.bat.String()
}

func RunBubble() {
	p := tea.NewProgram(NewModel())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}
