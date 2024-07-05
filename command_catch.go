package main

import "fmt"

func commandCatch(cfg *config, pokemon string) error {
	response, err := cfg.pokeapi.Catch(pokemon)
	if err != nil {
		return err
	}

	fmt.Println(response.Name)
	return nil
}
