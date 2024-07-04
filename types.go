package main

type pokedex_location_area struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}

type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
