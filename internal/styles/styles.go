package styles

import "charm.land/lipgloss/v2"

var (
// white     = lipgloss.Color("#FFFFFF")
// darkWhite = lipgloss.Color("#A7A7A7")
// green     = lipgloss.Color("#15E7CE")
// purple = lipgloss.Color("#855EFB")
)

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

// HTTP service status style
//var serviceServeStyle = lipgloss.NewStyle().
//	Bold(true).
//	Foreground((white))

//var (
//	TreeRootStyle = lipgloss.NewStyle().
//			Bold(true).
//			Foreground((white))

//	TreeChildStyle = lipgloss.NewStyle().
//			Bold(false).
//			Foreground((purple))
