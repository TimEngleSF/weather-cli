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
	if m.locSelection == 0 && m.location.zipcode == "" {
		return m.renderZipInput()
	}

	return m.unitSelection
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

		lipgloss.JoinVertical(lipgloss.Left, un, lipgloss.NewStyle().Bold(true).Render(qs), m.cBorderStyle.Render(cb.String())),
	)
}

func (m *model) renderZipInput() string {
	in := m.location.input
	inputField := in.View() // Render the input field

	// Define the style for the input field text
	inputFieldStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("12")) // Set the foreground color

	// Combine "Zipcode: " with the styled input field
	displayText := inputFieldStyle.Render(inputField)

	// Use the border style to render the combined text and place it
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		m.cBorderStyle.Render(displayText), // Render the combined text with the border style
	)
}
