package pokeapi

import (
	"encoding/json"
	"io"
	"log"
)

func (c *Client) GetMapLists(pageUrl *string) (ResponseFromPokedex, error) {
	pokedexUrl := url + "location-area"

	if pageUrl != nil {
		pokedexUrl = *pageUrl
	}

	res, err := c.httpClient.Get(pokedexUrl)
	if err != nil {
		return ResponseFromPokedex{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and message: %s", res.StatusCode, res.Body)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ResponseFromPokedex{}, err
	}

	var response ResponseFromPokedex
	if err := json.Unmarshal(body, &response); err != nil {
		return ResponseFromPokedex{}, err
	}

	return response, nil
}
