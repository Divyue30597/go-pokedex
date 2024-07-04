package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func commandGetPreviousMap() error {
	if requestVal >= 1 {
		requestVal -= 2
	}
	urlVal := fmt.Sprintf("location-area?offset=%v?limit=20", requestVal*20)
	res, err := http.Get(url + urlVal)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and message: %v", res.StatusCode, res.Body)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var locationArea pokedex_location_area
	if err := json.Unmarshal(body, &locationArea); err != nil {
		return err
	}

	for _, result := range locationArea.Results {
		fmt.Println(result.Name)
	}

	return nil
}
