package main

import (
	"bufio"
	"fmt"
	"internal/pokedexapi"
	"os"
	"strings"
)

type config struct {
	Next   string
	Prev   string
	client *pokedexapi.Client
}

func runRepl() {
	client := pokedexapi.PokedexClient()
	state := config{
		Next:   "",
		Prev:   "",
		client: client,
	}
	scanner := bufio.NewScanner(os.Stdin)
	commands := initCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		userInputArray := cleanInput(userInput)
		if len(userInputArray) > 0 {
			if command, ok := commands[userInputArray[0]]; ok {
				err := command.callback(&state)
				if err != nil {
					fmt.Println("Invalid command usage")
				}
			} else {
				fmt.Println("Command not found, use 'help' for a list of available commands")
			}
		} else {
			fmt.Println("No valid command found use help for getting help")
		}
	}
}

func cleanInput(text string) []string {
	text = strings.Trim(text, " ")
	splitStrings := strings.Split(text, " ")
	for index, str := range splitStrings {
		splitStrings[index] = strings.ToLower(strings.Trim(str, " "))
	}
	return splitStrings
}
