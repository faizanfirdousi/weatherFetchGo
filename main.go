package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Weather struct{
	Location struct{
		Name string `json:"name`
		Country string `json:"country"`
	} `json:"location"`

	Current struct{
		TempC float64 `json:"temp_c"`
		Condition struct{
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`

	Forecast struct{
		Forecastday [] struct{
			Hour []struct{
				TimeEpoch int64 `json:"time_epoch"`
				TempC float64 `json:"temp_c"`
				Condition struct{
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday`
	} `json:"forecast"`
}


func main() {
	// Load environment variables from .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		log.Fatal("WEATHER_API_KEY not set in environment")
	}

	// Construct the full URL
	url := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=18.584320985033493,73.73586345409191&days=1&aqi=no&alerts=no", apiKey)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Weather API not available: Status %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var weather Weather
	err = json.Unmarshal(body,&weather)
	if err!=nil {
		panic(err)
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	fmt.Printf(
		"%s, %s: %.0fC, %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.Condition.Text,
	)

	for _,hour := range hours{
		date := time.Unix(hour.TimeEpoch,0)

		if date.Before(time.Now()){
			continue
		}
		fmt.Printf(
			"%s - %.0fC,  %0.f%%,  %s\n",
			date.Format("15:04"),
			hour.TempC,
			hour.ChanceOfRain,
			hour.Condition.Text,
		)
	}

}
