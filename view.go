package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

/* VIEW */
func (m *model) View() string {
	if m.width == 0 {
		return "loading..."
	}
	/* ZIP CITY */
	if !m.initialSelected {
		var cb strings.Builder
		qs := "How would you like to search for the weather?\n"
		cb.WriteString("")
		for i, choice := range m.choices {
			line := "  " + choice
			if m.cursor == i {
				if i == 0 {
					line = "> " + m.highlightStyle.Render(choice)
				} else {
					line = "> " + m.highlightStyle.Render(choice)
				}
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

	if m.initialSelection == 0 && m.location.zipcode == "" {
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

	return m.location.zipcode
}
