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
	Zipcode       string `json:"zip"`
	City          string `json:"name"`
	State         string
	Country       string  `json:"country"`
	Lat           float64 `json:"lat"`
	Lon           float64 `json:"lon"`
	ZipInput      textinput.Model
	// CityInput     textinput.Model
	InputStyle    lipgloss.Style
	InputComplete bool
}

func New() *model {
	highlightStyle := HighlightStyle()
	cBorderStyle := ChoiceBorderStyle()
	unitStyle := UnitStyle()

	zipInput := textinput.New()
	zipInput.Prompt = lipgloss.NewStyle().Foreground(lipgloss.NoColor{}).Bold(true).Render("Zipcode: ")
	zipInput.Placeholder = "Enter Zipcode"
	zipInput.Focus()

	// cInput := textinput.New()
	// cInput.Prompt = lipgloss.NewStyle().Foreground(lipgloss.NoColor{}).Bold(true).Render("City: ")
	// cInput.Placeholder = "Enter City"

	u := helpers.GetUnits()
	if u != "" && !helpers.ValidateUnits(u) {
		u = ""
	}

	return &model{
		highlightStyle: highlightStyle,
		// locChoices:     []string{"zipcode", "city", "change units"},
		locChoices:     []string{"zipcode", "change units"},
		unitChoices:    []string{"imperial", "metric", "kelvin"},
		cBorderStyle:   cBorderStyle,
		// Location:       Location{ZipInput: zipInput, CityInput: cInput, InputComplete: false},
		Location:       Location{ZipInput: zipInput, InputComplete: false},
		unitSelection:  u,
		unitStyle:      unitStyle,
		resetUnit:      false,
		weatherData:    WeatherResponse{},
	}
}
