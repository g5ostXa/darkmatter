package styles

import "charm.land/lipgloss/v2"

var (
	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Align(lipgloss.Center).
			Foreground(lipgloss.Color("#FFFFFF")).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7b00ff"))

	TreeRootStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFFFFF"))

	TreeChildStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#6a5de3"))

	HttpChildStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#02BA84"))

	CommonStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#7b00ff"))

	LegendStyle = lipgloss.NewStyle().
			Bold(false).
			Foreground(lipgloss.Color("#A7A7A7"))
)
