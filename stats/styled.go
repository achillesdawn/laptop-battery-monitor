package stats

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
)

var tempStyle = lipgloss.NewStyle().Bold(true)

// var colors = []string{
// 	"51",
// 	"50",
// 	"49",
// 	"48",
// 	"47",
// 	"46",
// 	"82",
// 	"118",
// 	"154",
// 	"190",
// 	"226",
// 	"220",
// 	"214",
// 	"208",
// 	"202",
// 	"196",
// }

func ApplyColorScaleFloat(value float32) string {

	var color string

	switch {
	case value > 15:
		color = "196"
	case value > 14:
		color = "202"
	case value > 13:
		color = "208"
	case value > 12:
		color = "214"
	case value > 11:
		color = "220"
	case value > 10:
		color = "226"
	case value > 9:
		color = "154"
	case value > 8:
		color = "156"
	case value > 7:
		color = "85"
	case value > 6:
		color = "87"
	}

	return color
}

func FormatColor(color string, s string) string {
	return tempStyle.Foreground(lipgloss.Color(color)).Render(s)
}

func ColorFloat(value float32) string {
	color := ApplyColorScaleFloat(value)
	return FormatColor(color, fmt.Sprintf("%.1f", value))
}

func (b *BatStats) RenderStats() string {
	var s string

	color := ApplyColorScaleFloat(b.PowerNow)
	s = tempStyle.
		Foreground(lipgloss.Color(color)).
		Render(fmt.Sprintf("%.2f", b.PowerNow))

	return fmt.Sprintf(
		"%sw\t%s\t%.0f%%",
		s,
		b.TimeLeft.Round(time.Second*60).String(),
		b.Percent,
	)
}
