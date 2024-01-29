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
	weatherData    WeatherResponse
	Location       Location
	isLocSelected  bool
	locSelection   int
	width          int
	height         int
}

type Location struct {
	Zipcode    string `json:"zip"`
	City       string `json:"name"`
	State      string
	Country    string  `json:"country"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	Input      textinput.Model
	InputStyle lipgloss.Style
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
		Location:       Location{Input: input},
		unitSelection:  u,
		unitStyle:      unitStyle,
		resetUnit:      false,
		weatherData:    WeatherResponse{},
	}
}
