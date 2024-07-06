package main

import (
	"fmt"
)

func commandCatch(cfg *config, pokemon ...string) error {
	if len(pokemon) == 0 {
		return fmt.Errorf("no pokemon specified")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon[0])
	response, err := cfg.pokeapi.Catch(pokemon[0])
	if err != nil {
		return err
	}

	if _, ok := cfg.pokemonCaught.Get(pokemon[0]); ok {
		fmt.Println("Pokemon already caught")
		return nil
	}

	val := getRandomValue(response.BaseExperience)

	if val > response.BaseExperience {
		fmt.Printf("%s was caught!\n", pokemon[0])
		fmt.Println("you may now inspect it with inspect command.")
		cfg.pokemonCaught.Add(pokemon[0], response)
	} else {
		fmt.Printf("%s escaped!\n", pokemon[0])
	}

	return nil
}
