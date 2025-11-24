package main

type (
	Callback   func(*config, []string) error
	cliCommand struct {
		name        string
		description string
		callback    Callback
	}
)

var packageCommands map[string]cliCommand

func initCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Provides all the commands and their usage for the Pokedex",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Provides the next 20 location areas in Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "Provides the next 20 location areas in Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Takes a name of location and returns a list of Pokemon located there",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Takes a name of a pokemon and returns if you caught that pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Takes a name of a Pokemon and provides the stats if it exists in the Pokedex",
			callback:    commandInspect,
		},
	}
	packageCommands = commands
	return commands
}
