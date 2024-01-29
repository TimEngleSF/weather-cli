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
	highlightColor := lipgloss.Color("12")
	highlightStyle := lipgloss.NewStyle().Foreground(highlightColor)
	cBorderStyle := lipgloss.NewStyle().
		BorderForeground(lipgloss.Color("5")).
		BorderStyle(lipgloss.NormalBorder()).
		Width(80).Padding(0, 2).PaddingTop(1)

	input := textinput.New()
	input.Prompt = lipgloss.NewStyle().Foreground(lipgloss.NoColor{}).Bold(true).Render("Zipcode: ")
	input.Placeholder = "Enter Zipcode"
	input.Focus()

	u := helpers.GetUnits()

	return &model{
		highlightStyle: highlightStyle,
		locChoices:     []string{"zipcode", "city"},
		unitChoices:    []string{"imperial", "metric", "kelvin"},
		cBorderStyle:   cBorderStyle,
		location:       Location{input: input},
		unitSelection:  u,
	}
}
