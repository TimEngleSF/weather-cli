package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	cursor           int
	choices          []string
	cBorderStyle     lipgloss.Style
	highlightStyle   lipgloss.Style
	weatherData      []WeatherResponse
	location         Location
	done             bool
	initialSelected  bool
	initialSelection int
	width            int
	height           int
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
	input.Prompt = lipgloss.NewStyle().Foreground(lipgloss.NoColor{}).Bold(true).Render( "Zipcode: ")
	input.Placeholder = "Enter Zipcode"
	input.Focus()

	return &model{
		highlightStyle: highlightStyle,
		choices:        []string{"zipcode", "city"},
		cBorderStyle:   cBorderStyle,
		location:       Location{input: input},
	}
}