package main

import "github.com/Divyue30597/pokedex/internal/pokeapi"

type config struct {
	pokeapi          pokeapi.Client
	nextLocaleString *string
	prevLocalString  *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}
