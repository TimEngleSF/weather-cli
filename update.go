package main

import (
	"log"
	"weather-cli/helpers"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Main update function handling various states and inputs.
func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// Check if we need to reset to unit selection
	if m.resetUnit {
		m.unitSelection = ""
		m.isLocSelected = false
		// need to change locSelection, if it stays at selected position that triggers m.resetUnit to true, a loop will be created
		m.locSelection = 0
		m.cursor = 0
		m.resetUnit = false
	}

	// Proceed with unit selection if it's not yet set
	if m.unitSelection == "" {
		return m.handleUnitSelection(msg)
	}

	// Proceed with location selection if unit is selected but location is not
	if !m.isLocSelected {
		return m.handleLocationSelection(msg)
	}

	// Handle zipcode input last, assuming unit and location are selected
	return m.handleZipcodeInput(msg)
}

// Handles selction input for Units and Location Type
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
		}
	}
	return m, nil
}

// Handles unit selection logic.
func (m *model) handleUnitSelection(msg tea.Msg) (tea.Model, tea.Cmd) {
	updatedModel, cmd := m.processChoice("units", msg)
	helpers.WriteUnits(m.unitSelection) // Persist selected units
	return updatedModel, cmd
}

// Handles location selection logic based on chosen method (zipcode or city).
func (m *model) handleLocationSelection(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m.processChoice("loc", msg)

}

// Handles direct zipcode input for location.
func (m *model) handleZipcodeInput(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.locSelection == 0 && m.Location.Zipcode == "" {
		return m.processZipcodeInput(msg)
	}
	updatedModel, cmd := m.handleGlobalKeys(msg)
	m, ok := updatedModel.(*model)
	if !ok {
		log.Fatalf("Update: Unable to assert model type during '%s'", "handleZipcodeInput")
	}

	m.SetCurrWeatherByZip()
	return m, cmd
}

// Process user choice for either units or location.
func (m *model) processChoice(ct string, msg tea.Msg) (tea.Model, tea.Cmd) {
	updatedModel, cmd := m.SelectChoices(ct, msg)
	m, ok := updatedModel.(*model)
	// if change unit was selected during location reset, set m.resetUnit to true
	if ct == "loc" {
		if m.locSelection == len(m.locChoices)-1 {
			m.resetUnit = true
		}
	}
	if !ok {
		log.Fatalf("Update: Unable to assert model type for '%s'", ct)
	}
	return m, cmd
}

// Process direct input of zipcode.
func (m *model) processZipcodeInput(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	in := m.Location.Input

	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleZipcodeKey(msg, in)
	}

	return m, cmd
}

// Handle key inputs specifically for zipcode entry.
func (m *model) handleZipcodeKey(msg tea.KeyMsg, in textinput.Model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	ms := msg.String()
	val := in.Value()

	// Check if the key press is a digit before updating the input model
	if len(val) < 5 && isDigit(msg) || ms == "backspace" || ms == "ctrl+c" || ms == "enter" {
		in, cmd = in.Update(msg)
		m.Location.Input = in
	}

	if ms == "ctrl+c" {
		return m, tea.Quit
	}

	if len(val) == 5 && (ms == "enter" || ms == " ") {
		m.Location.Zipcode = in.Value()
		in.SetValue("")
	}

	return m, cmd
}

// Handle global key inputs like window size changes or quit command.
func (m *model) handleGlobalKeys(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	}
	return m, nil
}

// Utility function to check if a key press is a digit.
func isDigit(keyMsg tea.KeyMsg) bool {
	return keyMsg.Type == tea.KeyRunes && len(keyMsg.Runes) == 1 && keyMsg.Runes[0] >= '0' && keyMsg.Runes[0] <= '9'
}
