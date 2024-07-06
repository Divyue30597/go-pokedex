package main

import "fmt"

func commandExplore(cfg *config, location ...string) error {
	if len(location) == 0 {
		return fmt.Errorf("no location specified")
	}

	fmt.Printf("Exploring %s\n", location[0])
	fmt.Println("Found Pokemon:")

	response, err := cfg.pokeapi.Explore(location[0])
	if err != nil {
		return err
	}

	for _, pokemon := range response.PokemonEncounters {
		fmt.Println("- " + pokemon.Pokemon.Name)
	}

	return nil
}
