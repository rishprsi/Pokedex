package main

import (
	"fmt"
	"os"
)

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

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("")
	for _, command := range packageCommands {
		// use := command.name + ": " + command.description
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	return nil
}
