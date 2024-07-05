package pokeexplore

import (
	"encoding/json"
	"io"
	"net/http"
)

func (e ExploringLocation) Explore(baseUrl string, location string) (ExploringLocation, error) {
	newUrl := baseUrl + "location-area/" + location

	res, err := http.Get(newUrl)
	if err != nil {
		return ExploringLocation{}, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ExploringLocation{}, err
	}

	var response ExploringLocation
	if err := json.Unmarshal(body, &response); err != nil {
		return ExploringLocation{}, err
	}

	return response, nil
}
