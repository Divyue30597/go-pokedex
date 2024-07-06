package main

import "fmt"

func commandInspect(cfg *config, pokemon ...string) error {
	if len(pokemon) == 0 {
		return fmt.Errorf("no pokemon specified")
	}

	if val, ok := cfg.pokemonCaught.Get(pokemon[0]); ok {
		fmt.Println("Name: ", val.Name)
		fmt.Println("Height: ", val.Height)
		fmt.Println("Weight: ", val.Weight)
		fmt.Println("Stats:")
		for _, stat := range val.Stats {
			fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pokemon_type := range val.Types {
			fmt.Printf("  - %s\n", pokemon_type.Type.Name)
		}

	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}
