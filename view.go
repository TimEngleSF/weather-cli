package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

/* VIEW */
func (m *model) View() string {
	if m.width == 0 {
		return "loading..."
	}
	/* UNIT SELECTION */
	if m.unitSelection == "" || m.resetUnit {
		return m.renderUnitSelection()
	}
	/* LOCATION TYPE SELECTION */
	if !m.isLocSelected {
		return m.renderLocSelection()
	}
	/* ZIPCODE INPUT */
	if m.locSelection == 0 && m.Location.Zipcode == "" {
		return m.renderZipInput()
	}

	/* WEATHER FORECAST */
	if m.weatherData.Name == "" {
		return m.FetchWeatherDisplay()
	}

	return fmt.Sprintf("%+v", m.weatherData)
}

func (m *model) renderUnitSelection() string {
	var cb strings.Builder
	qs := "Which units would you like the weather displayed in?\n"
	cb.WriteString("")
	for i, choice := range m.unitChoices {
		line := "  " + choice
		if m.cursor == i {
			line = "> " + m.highlightStyle.Render(choice)
		}
		cb.WriteString(line + "\n")
	}

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,

		lipgloss.JoinVertical(lipgloss.Left, lipgloss.NewStyle().Bold(true).Render(qs), m.cBorderStyle.Render(cb.String())),
	)
}

func (m *model) renderLocSelection() string {
	var cb strings.Builder
	un := fmt.Sprintf("Units: %s\n", m.unitStyle.Render(m.unitSelection))
	qs := "How would you like to search for the weather?\n"
	cb.WriteString("")
	for i, choice := range m.locChoices {
		line := "  " + choice
		if m.cursor == i {
			line = "> " + m.highlightStyle.Render(choice)
		}
		cb.WriteString(line + "\n")
	}

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,

		lipgloss.JoinVertical(
			lipgloss.Left, 
			un, 
			lipgloss.NewStyle().Bold(true).Render(qs), 
			m.cBorderStyle.Render(cb.String()),
		),
	)
}

func (m *model) renderZipInput() string {
	in := m.Location.Input
	inputField := in.View() 

	inputFieldStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("12"))

	displayText := inputFieldStyle.Render(inputField)

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		m.cBorderStyle.Render(displayText),
	)
}
