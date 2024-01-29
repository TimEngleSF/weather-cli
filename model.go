package main

import (
	"weather-cli/helpers"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	cursor         int
	locChoices     []string
	unitSelection  string
	unitChoices    []string
	unitStyle      lipgloss.Style
	resetUnit      bool
	cBorderStyle   lipgloss.Style
	highlightStyle lipgloss.Style
	weatherData    []WeatherResponse
	location       Location
	done           bool
	isLocSelected  bool
	locSelection   int
	width          int
	height         int
}

type Location struct {
	zipcode    string
	city       string
	state      string
	country    string
	input      textinput.Model
	inputStyle lipgloss.Style
}

func New() *model {
	highlightStyle := HighlightStyle()
	cBorderStyle := ChoiceBorderStyle()
	unitStyle := UnitStyle()

	input := textinput.New()
	input.Prompt = lipgloss.NewStyle().Foreground(lipgloss.NoColor{}).Bold(true).Render("Zipcode: ")
	input.Placeholder = "Enter Zipcode"
	input.Focus()

	u := helpers.GetUnits()
	if u != "" && !helpers.ValidateUnits(u) {
		u = ""
	}

	return &model{
		highlightStyle: highlightStyle,
		locChoices:     []string{"zipcode", "city", "change units"},
		unitChoices:    []string{"imperial", "metric", "kelvin"},
		cBorderStyle:   cBorderStyle,
		location:       Location{input: input},
		unitSelection:  u,
		unitStyle:      unitStyle,
	}
}
