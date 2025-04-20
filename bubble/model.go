package bubble

import (
	"fmt"
	"math"
	"time"

	"github.com/achillesdawn/battery-monitor/stats"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	bat  *stats.BatStats
	last []float32
}

func NewModel() Model {
	bat, err := stats.New()
	if err != nil {
		panic(err)
	}

	bat.CalcTimeLeft()

	last := make([]float32, 0, numSamples)

	return Model{
		bat:  bat,
		last: last,
	}
}

func (m Model) LastStats() string {
	var minVal float32 = float32(math.Inf(1))
	var maxVal float32 = 0
	var sumVal float32 = 0

	for _, val := range m.last {
		if val > maxVal {
			maxVal = val
		} else if val < minVal {
			minVal = val
		}

		sumVal += val
	}

	avgVal := sumVal / float32(len(m.last))

	return fmt.Sprintf("%.1fw  %.1fw %.1fw", minVal, avgVal, maxVal)

}

func (m *Model) Monitor(t time.Time) tea.Msg {
	m.bat.ReadPowerAndEnergy()

	m.bat.CalcTimeLeft()

	return Message{}
}
