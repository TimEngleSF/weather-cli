package main

import "github.com/charmbracelet/lipgloss"

func ChoiceBorderStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		BorderForeground(lipgloss.Color("5")).
		BorderStyle(lipgloss.NormalBorder()).
		Width(80).Padding(0, 2).PaddingTop(1)
}

func HighlightStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
}

func UnitStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
}

func (m *model) FetchWeatherDisplay() string {
	hc := lipgloss.NewStyle().Foreground(lipgloss.Color("8"))

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		hc.Render("Fetching Weather..."),
	)
}
