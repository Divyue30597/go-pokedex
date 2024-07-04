package main

import (
	"time"

	"github.com/Divyue30597/pokedex/pokeapi"
)

var requestVal int = 0

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := config{
		pokeapi: pokeClient,
	}

	startRepl(&cfg)
}
