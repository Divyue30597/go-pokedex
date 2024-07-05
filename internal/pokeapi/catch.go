package pokeapi

import (
	"encoding/json"
	"io"
	"log"
)

func (c *Client) Catch(pokemon string) (Pokemon, error) {
	url := baseUrl + "pokemon/" + pokemon

	res, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and message: %s", res.StatusCode, res.Body)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	var response Pokemon
	if err := json.Unmarshal(body, &response); err != nil {
		return Pokemon{}, err
	}

	return response, nil
}
