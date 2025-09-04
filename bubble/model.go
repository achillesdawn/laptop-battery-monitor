package bubble

import (
	"fmt"
	"math"
	"time"

	"github.com/achillesdawn/laptop-battery-monitor/stats"
	tea "github.com/charmbracelet/bubbletea"
)

type App struct {
	bat  *stats.BatStats
	last []float32

	width  int
	height int

	quit bool
}

func NewApp() App {
	bat, err := stats.New()
	if err != nil {
		panic(err)
	}

	_, err = bat.CalcTimeLeft()
	if err != nil {
		panic(err)
	}

	last := make([]float32, 0, numSamples)

	return App{
		bat:  bat,
		last: last,
		quit: false,
	}
}

func (m App) LastStats() string {
	var minVal = float32(math.Inf(1))
	var maxVal float32 = 0
	var sumVals float32 = 0

	for _, val := range m.last {
		if val > maxVal {
			maxVal = val
		}
		if val < minVal {
			minVal = val
		}

		sumVals += val
	}

	avgVal := sumVals / float32(len(m.last))

	minValStr := stats.ColorFloat(minVal)
	avgValStr := stats.ColorFloat(avgVal)
	maxValStr := stats.ColorFloat(maxVal)

	return fmt.Sprintf("%sw %sw %sw", minValStr, avgValStr, maxValStr)
}

func (m *App) Monitor(t time.Time) tea.Msg {
	err := m.bat.ReadPowerAndEnergy()
	if err != nil {
		panic(err)
	}

	_, err = m.bat.CalcTimeLeft()
	if err != nil {
		panic(err)
	}

	return AppMessage{}
}
