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
	var minV float32 = float32(math.Inf(1))
	var maxV float32 = 0
	var sumV float32 = 0

	for _, val := range m.last {
		if val > maxV {
			maxV = val
		} else if val < minV {
			minV = val
		}

		sumV += val
	}

	avgVal := sumV / float32(len(m.last))

	return fmt.Sprintf("%.1fw  %.1fw %.1fw", minV, avgVal, maxV)

}

func (m *Model) Monitor(t time.Time) tea.Msg {
	m.bat.ReadPowerAndEnergy()

	m.bat.CalcTimeLeft()

	return Message{}
}
