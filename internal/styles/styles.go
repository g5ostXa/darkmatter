package styles

import "charm.land/lipgloss/v2"

var HeaderStyle = lipgloss.NewStyle().
	Bold(true).
	Align(lipgloss.Center).
	Foreground(lipgloss.Color("#FFFFFF")).
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#5d5de3"))

var TreeRootStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FFFFFF"))

var TreeChildStyle = lipgloss.NewStyle().
	Bold(false).
	Foreground(lipgloss.Color("#5edcfb"))

var CommonStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#5d5de3"))

var (
	white     = lipgloss.Color("#FFFFFF")
	darkWhite = lipgloss.Color("#A7A7A7")
	//green     = lipgloss.Color("#15E7CE")
	purple = lipgloss.Color("#855EFB")
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

//	LegendStyle = lipgloss.NewStyle().
//			Bold(false).
//			Foreground((darkWhite))
//)
