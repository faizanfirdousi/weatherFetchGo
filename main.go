package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

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

	fmt.Println(string(body))
}
