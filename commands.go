package main

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exits the program",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in pokemon world",
			callback:    commandGetMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of previous 20 location areas in pokemon world",
			callback:    commandGetPreviousMap,
		},
	}
}
