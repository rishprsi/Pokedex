package main

import (
	"fmt"
	"os"
)

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
