package main

type (
	Callback   func(*config) error
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
	}
	packageCommands = commands
	return commands
}
