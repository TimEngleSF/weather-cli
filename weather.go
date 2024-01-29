package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/carlmjohnson/requests"
	"github.com/charmbracelet/lipgloss"
	"github.com/joho/godotenv"
)

type WeatherResponse struct {
	Temp struct {
		Current float64 `json:"temp"`
		Min     float64 `json:"temp_min"`
		Max     float64 `json:"temp_max"`
	} `json:"main"`
	Weather []struct {
		Id          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Name  string `json:"name"`
	Time  time.Time
	Style lipgloss.Style
}

func (m *model) SetLocation() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	api_key := os.Getenv("API_KEY")

	var url string
	if len(m.Location.Zipcode) == 5 {
		url = fmt.Sprintf("http://api.openweathermap.org/geo/1.0/zip?zip=%s,%s&appid=%s", m.Location.Zipcode, "us", api_key)
	}

	err = requests.URL(url).ToJSON(&m.Location).Fetch(context.TODO())

	if err != nil {
		log.Println("SetLocation: ", err)
	}
}

func (m *model) SetCurrWeatherByZip() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	api_key := os.Getenv("API_KEY")

	m.SetLocation()

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?zip=%s,%s&appid=%s&units=%s", m.Location.Zipcode, "us", api_key, m.unitSelection)

	var resp WeatherResponse
	err = requests.URL(url).ToJSON(&resp).Fetch(context.TODO())
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println("could not connect to jsonplaceholder.typicode.com:", err)
	}
	resp.Time = time.Now()
	m.weatherData = resp
}
