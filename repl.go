package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Divyue30597/pokedex/internal/pokeapi"
)

type config struct {
	pokeapi          pokeapi.Client
	nextLocaleString *string
	prevLocalString  *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
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

		if ok {
			if err := command.callback(cfg); err != nil {
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
