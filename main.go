package main

import (
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	cursor           int
	choices          []string
	cBorderStyle     lipgloss.Style
	highlightStyle   lipgloss.Style
	weatherData      []WeatherData
	location         Location
	done             bool
	initialSelected  bool
	initialSelection int
	width            int
	height           int
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
	cBorderStyle := lipgloss.NewStyle().
	BorderForeground(lipgloss.Color("5")).
	BorderStyle(lipgloss.NormalBorder()).
	Width(80).Padding(0, 2).PaddingTop(1)

	return &model{
		highlightStyle: highlightColor,
		choices:        []string{"zipcode", "city"},
		cBorderStyle: cBorderStyle,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

/* UPDATE */
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	/* ZIP CITY PROMPT */
	if !m.initialSelected {
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

	/* DEFAULT */
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

	/* ZIPCODE INPUT */
	// if m.initialSelected && m.initialSelected == 0{
	// 	switch msg := msg.(type){}
	// }

	return m, nil
}

/* VIEW */
func (m model) View() string {
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

			lipgloss.JoinVertical(lipgloss.Left, qs, m.cBorderStyle.Render(cb.String())),
		)
	}
	return m.choices[m.initialSelection]
}

/* MAIN */
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
