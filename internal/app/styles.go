package app

import "github.com/charmbracelet/lipgloss"

var (
	violet     = lipgloss.Color("#a78bfa")
	violetDark = lipgloss.Color("#7c3aed")
	violetDim  = lipgloss.Color("#8b5cf6")
)

var (
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(violet).
			Padding(0, 1)

	ItemStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#e0e0e0")).
			Padding(0, 1)

	SelectedItemStyle = lipgloss.NewStyle().
				Background(violetDark).
				Foreground(lipgloss.Color("#ffffff")).
				Padding(0, 1)

	FooterStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6b7280"))

	LoadingStyle = lipgloss.NewStyle().
			Foreground(violet).
			Italic(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ef4444")).
			Bold(true)

	docStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(violet).
			Padding(1, 2)
)
