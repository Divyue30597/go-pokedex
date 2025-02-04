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
			description: "Get next page of Pokemon locations (limit is 20)",
			callback:    commandGetMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get previous page of Pokemon locations (limit is 20)",
			callback:    commandGetPreviousMap,
		},
		"explore": {
			name:        "explore",
			description: "Let's you explore the location from the selected lists.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a caught pokemon.",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Get list of caught pokemon",
			callback:    commandPokedex,
		},
	}
}
