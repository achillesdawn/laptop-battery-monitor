package bubble

import "github.com/charmbracelet/lipgloss"

var BorderedStyle = lipgloss.
	NewStyle().
	// Padding(1).
	Align(lipgloss.Left).
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("228"))
