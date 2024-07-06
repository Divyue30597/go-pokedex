package main

import "fmt"

func commandPokedex(cfg *config, _ ...string) error {
	fmt.Println("Your Pokedex:")
	val, err := cfg.pokemonCaught.GetAll()
	if err != nil {
		return err
	}

	for _, pokemon := range val {
		fmt.Println("  - " + pokemon)
	}

	return nil
}
