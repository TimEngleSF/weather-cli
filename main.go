package main

import (
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	cursor         int
	choices        []string
	highlightStyle lipgloss.Style
	weatherData     []WeatherData
	location        Location
	done            bool
	initialSelected bool
	initialSelection int
	width           int
	height          int
}

type WeatherData struct {
	temp      float64
	date      string
	condition string
	high      float64
	low       float64
}

type Location struct {
	zipcode string
	city    string
	state   string
	country string
}

func New() *model {
	highlightColor := lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
	return &model{
		highlightStyle: highlightColor,
		choices: []string{"zipcode", "city"},
	}
}



func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if !m.initialSelected{
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
			if m.cursor < len(m.choices) {
				m.cursor++
			}
		case "enter", "space":
			m.initialSelection = m.cursor
			m.initialSelected = true
		}
	}
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		switch msg.String(){
		case "ctrl+c":
			return m, tea.Quit
		}
	}


	return m, nil
}

func (m model) View() string {
	if m.width == 0 {
		return "loading..."
	}

	if !m.initialSelected {
		var sb strings.Builder
		sb.WriteString("How would you like to search for the weather?\n\n")

		for i, choice := range m.choices {
			line := "  " + choice
			if m.cursor == i {
				line = "> " + m.highlightStyle.Render(choice)
			}
			sb.WriteString(line + "\n")
		}

		return sb.String()
	}
	return m.choices[m.initialSelection]
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}
	// Close the log file when main finishes executing
	defer f.Close()
	initialModel := New()
	p := tea.NewProgram(initialModel, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
