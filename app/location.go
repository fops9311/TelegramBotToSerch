// location.go

package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func getLocation(local, API string) ([]string, error) {
	var builder strings.Builder
	builder.WriteString("https://hotels4.p.rapidapi.com/locations/v3/search?q=")
	builder.WriteString(local)
	builder.WriteString("&locale=ru_RU&langid=1031&siteid=300000001")

	url := builder.String()

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", API)
	req.Header.Add("X-RapidAPI-Host", "hotels4.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []string{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	body, _ := io.ReadAll(res.Body)

	jsonResponse := string(body)
	var locationResponse LocationResponse
	err = json.Unmarshal([]byte(jsonResponse), &locationResponse)
	if err != nil {
		log.Println("Error decoding JSON:", err)
		return []string{}, err
	}

	var cities []string

	for _, result := range locationResponse.SR {
		if result.Type == "CITY" {
			cities = append(cities, result.RegionNames.ShortName)
		}
	}

	return cities, nil
}
