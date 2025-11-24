package commands

import (
	"pokecache"
	"pokedexapi"
	"time"
)

type (
	Callback   func(*config) error
	cliCommand struct {
		Name        string
		Description string
		Callback    Callback
	}
	config struct {
		Next   string
		Prev   string
		client *pokedexapi.Client
		cache  *pokecache.Cache
	}
)

var packageCommands map[string]cliCommand

func InitCommands() (map[string]cliCommand, *config) {
	client := pokedexapi.PokedexClient()
	cache := pokecache.NewCache(time.Duration(60 * time.Second))
	state := config{
		Next:   "",
		Prev:   "",
		client: client,
		cache:  cache,
	}
	commands := map[string]cliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Provides all the commands and their usage for the Pokedex",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Provides the next 20 location areas in Pokemon world",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "map",
			Description: "Provides the next 20 location areas in Pokemon world",
			Callback:    commandMapb,
		},
	}
	packageCommands = commands
	return commands, &state
}
