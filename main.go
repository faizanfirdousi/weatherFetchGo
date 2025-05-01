package main

import (
	"net/http"
	"io"
	"fmt"
)



func main(){
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=994231216f7a4852841212932250105&q=18.584320985033493,73.73586345409191&days=1&aqi=no&alerts=no")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil{
		panic(err)
	}

	fmt.Println(string(body))
}
