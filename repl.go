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
