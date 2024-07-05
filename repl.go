package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Divyue30597/pokedex/internal/pokeapi"
	"github.com/Divyue30597/pokedex/internal/pokeexplore"
)

type config struct {
	pokeapi          pokeapi.Client
	pokemonList      pokeexplore.ExploringLocation
	nextLocaleString *string
	prevLocalString  *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := cleanText(scanner.Text())
		command, ok := getCommands()[input[0]]

		var locationVal string = ""

		if len(input) >= 2 {
			locationVal = input[1]
		}

		if ok {
			if err := command.callback(cfg, locationVal); err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Command not found")
			continue
		}
	}
}

func cleanText(text string) []string {
	newText := strings.ToLower(text)
	finalText := strings.Fields(newText)
	return finalText
}
