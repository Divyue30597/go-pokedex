package main

import (
	"time"

	"github.com/Divyue30597/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := config{
		pokeapi: pokeClient,
	}

	startRepl(&cfg)
}
