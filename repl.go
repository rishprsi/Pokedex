package main

import (
	"bufio"
	"fmt"
	"internal/pokecache"
	"internal/pokedexapi"
	"os"
	"strings"
	"time"
)

type config struct {
	Next    string
	Prev    string
	BaseURL string
	Page    int
	client  *pokedexapi.Client
	cache   *pokecache.Cache
	Pokedex map[string]Pokemon
}

func runRepl() {
	client := pokedexapi.PokedexClient()
	cache := pokecache.NewCache(time.Duration(60 * time.Second))
	state := config{
		Next:    "",
		Prev:    "",
		BaseURL: "https://pokeapi.co/api/v2/",
		Page:    0,
		client:  client,
		cache:   cache,
		Pokedex: map[string]Pokemon{},
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
				err := command.callback(&state, userInputArray[1:])
				if err != nil {
					fmt.Println("Invalid command usage")
					fmt.Println(err)
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
