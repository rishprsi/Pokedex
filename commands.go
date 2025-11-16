package main

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

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
	}
	packageCommands = commands
	return commands
}
