package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func ChoiceBorderStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		BorderForeground(lipgloss.Color("5")).
		BorderStyle(lipgloss.NormalBorder()).
		Width(80).Padding(0, 2).PaddingTop(1)
}

func HighlightStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
}

func UnitStyle() lipgloss.Style {
	return lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
}

func (m *model) FetchWeatherDisplay() string {
	hc := lipgloss.NewStyle().Foreground(lipgloss.Color("8"))

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		hc.Render("Fetching Weather..."),
	)
}

func (m *model) WeatherCard() string {
	t := m.weatherData.Temp
	w := m.weatherData.Weather[0]
	n := m.weatherData.Name

	var cBgColor string
	var cFgColor string
	var mFgColor string
	wId := w.Id
	switch {
	case wId == 800:
		cBgColor = "#01BFFF"
		cFgColor = "#F1F1F1"
		mFgColor = "#FFCD00"
	case wId > 800 && wId < 900:
		cBgColor = "#7F9BA6"
		cFgColor = "#F1F1F1"
		mFgColor = "#FFFAE0"
	case wId > 600 && wId < 700:
		cBgColor = "#FFF"
		cFgColor = "#000"
		mFgColor = "#1D71F2"
	case wId > 300 && wId < 600:
		cBgColor = "#7F9BA6"
		cFgColor = "#F1F1F1"
		mFgColor = "#1D71F2"
	case wId > 200 && wId < 300:
		cBgColor = "#4E6969"
		cFgColor = "#F1F1F1"
		mFgColor = "#DC2626"
	}

	cardWidth := 50

	lineStyle := lipgloss.NewStyle().
		Width(cardWidth - 8). // Adjust for padding
		PaddingLeft(4).       // Match card padding
		PaddingRight(4).
		Bold(true).
		Background(lipgloss.Color(cBgColor)).
		Foreground(lipgloss.Color(cFgColor)).
		Align(lipgloss.Left)

	nStr := lineStyle.Render(n)
	tCStr := lineStyle.Copy().PaddingBottom(1).Render(fmt.Sprintf("Current: %v", t.Current))
	tLStr := lineStyle.Render(fmt.Sprintf("Low: %v", t.Min))
	tHStr := lineStyle.Render(fmt.Sprintf("High: %v", t.Max))
	wDescStr := lineStyle.Copy().PaddingBottom(1).Render(w.Description)
	wMainStr := lineStyle.Copy().Foreground(lipgloss.Color(mFgColor)).Render(w.Main)

	cardContent := lipgloss.JoinVertical(lipgloss.Left, nStr, tCStr, wMainStr, wDescStr, tHStr, tLStr)

	cardStyle := lipgloss.NewStyle().
		Width(cardWidth).
		Background(lipgloss.Color(cBgColor)).
		PaddingLeft(4).
		PaddingRight(4).
		PaddingTop(1).
		PaddingBottom(1).
		Align(lipgloss.Center).MarginBottom(2)

	card := cardStyle.Render(cardContent)

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Center, card, "Press a key to find more forecasts"),
	)
}
