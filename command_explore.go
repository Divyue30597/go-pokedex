package main

import "fmt"

func commandExplore(cfg *config, location string) error {
	fmt.Printf("Exploring %s\n", location)
	fmt.Println("Found Pokemon:")

	response, err := cfg.pokemonList.Explore(baseUrl, location)
	if err != nil {
		return err
	}

	for _, pokemon := range response.PokemonEncounters {
		fmt.Println("- " + pokemon.Pokemon.Name)
	}

	return nil
}
