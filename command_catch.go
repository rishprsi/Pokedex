package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
)

type Pokemon struct {
	ID             int              `json:"id"`
	Name           string           `json:"name"`
	BaseExperience int              `json:"base_experience"`
	Height         int              `json:"height"`
	Order          int              `json:"order"`
	Weight         int              `json:"weight"`
	Abilities      []map[string]any `json:"abilties"`
	Stats          []PokeStat       `json:"stats"`
	Types          []PokeType       `json:"types"`
}

type PokeStat struct {
	BaseStat int        `json:"base_stat"`
	Stat     NameStruct `json:"stat"`
}

type PokeType struct {
	Type NameStruct `json:"type"`
}

type NameStruct struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func commandCatch(state *config, args []string) error {
	if len(args) == 0 {
		return errors.New("usage of 'catch' command: catch pokemon_name")
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", args[0])
	url := fmt.Sprintf("%spokemon/%s/", state.BaseURL, args[0])
	body, err := cacheHit(state, url)
	if err != nil {
		return fmt.Errorf("failed to fetch Pokemon details with the following error:\n %v", err)
	}
	pokemon, err := byteToPokemon(body)
	if err != nil {
		return err
	}
	caught := catchEmAll(pokemon.BaseExperience)
	if caught {
		state.Pokedex[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}
	return nil
}

func commandInspect(state *config, args []string) error {
	if len(args) == 0 {
		return errors.New("usage of 'catch' command: catch pokemon_name")
	}
	pokemon, ok := state.Pokedex[args[0]]
	if !ok {
		fmt.Printf("%s is not in your Pokedex, catchEm first\n", args[0])
	} else {
		printPokemon(pokemon)
	}
	return nil
}

func byteToPokemon(body []byte) (Pokemon, error) {
	var pokemon Pokemon
	err := json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error unpacking pokemon with errer:\n%v", err)
	}
	return pokemon, nil
}

func catchEmAll(baseExperience int) bool {
	minProb, maxProb := .1, .9
	random := rand.Float64()
	prob := float64(baseExperience) / 600.0
	if prob > maxProb {
		prob = maxProb
	} else if prob < minProb {
		prob = minProb
	}
	fmt.Printf("Required: %v, Actual %v\n", baseExperience, random*600)
	if random >= prob {
		return true
	} else {
		return false
	}
}

func printPokemon(pokemon Pokemon) {
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  --%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("  -- %s\n", typ.Type.Name)
	}
}
