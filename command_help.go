package main

import "fmt"

func commandHelp(cfg *config, _ ...string) error {
	fmt.Println("Welcome to the Pokedex:")
	fmt.Println("Usage:")
	fmt.Println()

	commands := getCommands()
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	fmt.Println()
	return nil
}
