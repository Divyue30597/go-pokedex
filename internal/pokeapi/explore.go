package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) Explore(location string) (ExploringLocation, error) {
	newUrl := baseUrl + "location-area/" + location

	// get from cache
	if val, ok := c.cache.Get(newUrl); ok {
		locationArea := ExploringLocation{}

		if err := json.Unmarshal(val, &locationArea); err != nil {
			return ExploringLocation{}, err
		}

		return locationArea, nil
	}

	res, err := http.Get(newUrl)
	if err != nil {
		return ExploringLocation{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and message: %s", res.StatusCode, res.Body)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ExploringLocation{}, err
	}

	var response ExploringLocation
	if err := json.Unmarshal(body, &response); err != nil {
		return ExploringLocation{}, err
	}

	// add to cache
	c.cache.Add(newUrl, body)
	return response, nil
}
