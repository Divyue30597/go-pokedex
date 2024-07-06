package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := cleanText(scanner.Text())
		command, ok := getCommands()[input[0]]

		args := []string{}

		if len(input) > 1 {
			args = input[1:]
		}

		if ok {
			if err := command.callback(cfg, args...); err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Command not found")
			continue
		}
	}
}
