package main

import (
	"log"
	"weather-cli/helpers"

	tea "github.com/charmbracelet/bubbletea"
)

/* UPDATE */
func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	/* SELECT UNIT TYPE */ 
	if m.unitSelection == ""{
		updatedModel, cmd := m.SelectChoices("units", msg)
		m, ok := updatedModel.(*model)
		if !ok {
			log.Fatal("Update:  Unable to assert model")
		}
		return m, cmd
	}

	/* SEARCH BY LOC TYPE SELECTION */ 
	if !m.isLocSelected {
		updatedModel, cmd := m.SelectChoices("loc", msg)
		m, ok := updatedModel.(*model)
		if !ok {
			log.Fatal("Update:  Unable to assert model")
		}
		helpers.WriteUnits(m.unitSelection)
		return m, cmd
	} else if m.locSelection == 0 && m.location.zipcode == "" {
		/* ZIPCODE INPUT */
		in := m.location.input
		var cmd tea.Cmd

		switch msg := msg.(type) {
		case tea.KeyMsg:
			ms := msg.String()
			val := in.Value()
			// Check if the key press is a digit before updating the input model
			if len(val) < 5 && isDigit(msg) || ms == "backspace" || ms == "ctrl+c" || ms == "enter" {
				in, cmd = in.Update(msg)
				m.location.input = in
			}

			if ms == "ctrl+c" {
				return m, tea.Quit
			}

			if len(val) == 5 && ms == "enter" || ms == " " {
				if len(in.Value()) == 5 {
					m.location.zipcode = in.Value()
					in.SetValue("") // Reset the input field after capturing the zipcode
				}
			}
		}
		return m, cmd

	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, cmd
}
func isDigit(keyMsg tea.KeyMsg) bool {
	return keyMsg.Type == tea.KeyRunes && len(keyMsg.Runes) == 1 && keyMsg.Runes[0] >= '0' && keyMsg.Runes[0] <= '9'
}

// Update function to choose display units and weather forecast location.
// ct should be "units" for unit selection.
// ct should be "loc" for location selection.
func (m *model) SelectChoices(ct string, msg tea.Msg) (tea.Model, tea.Cmd) {
	var choices []string

	if ct == "loc" {
		choices = m.locChoices
	} else if ct == "units" {
		choices = m.unitChoices
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(choices)-1 {
				m.cursor++
			}
		case "enter", "space":
			if ct == "loc" {
				m.isLocSelected = true
				m.locSelection = m.cursor
			} else if ct == "units" {
				m.unitSelection = m.unitChoices[m.cursor]
			}
			m.cursor = 0
			// Initialize or focus the text input here if needed
		}
	}
	return m, nil
}
