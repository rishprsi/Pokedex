package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type LocationDetails struct {
	GameIndex         int          `json:"game_index"`
	ID                int          `json:"id"`
	Location          LocationArea `json:"location"`
	Name              string       `json:"name"`
	PokemonEncounters []Encounter  `json:"pokemon_encounters"`
}

type Encounter struct {
	Pokemon        PokemonShort     `json:"pokemon"`
	VersionDetails []map[string]any `json:"version_details"`
}

type PokemonShort struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func commandExplore(state *config, args []string) error {
	fmt.Printf("Exploring %s...\n", args[0])
	if len(args) == 0 {
		return errors.New("invalid use of command explore, usage: 'explore location'")
	}
	url := fmt.Sprintf("%slocation-area/%s/", state.BaseURL, args[0])
	fmt.Println(url)
	body, err := cacheHit(state, url)
	if err != nil {
		return fmt.Errorf("encountered Error while retreiving Location Details: %v", err)
	}
	locationDetails, err := byteToLocationDetails(body)
	if err != nil {
		return err
	}
	err = locationToPokemons(locationDetails)
	if err != nil {
		return fmt.Errorf("encountered error while unpacking pokemon from location details: %v", err)
	}
	return nil
}

func locationToPokemons(locationDetails LocationDetails) error {
	if locationDetails.Name == "" || locationDetails.PokemonEncounters == nil {
		fmt.Println(locationDetails)
		return errors.New("no pokemon found in the location")
	}
	fmt.Println("Found Pokemon:")
	for _, value := range locationDetails.PokemonEncounters {
		fmt.Printf("-- %s\n", value.Pokemon.Name)
	}
	return nil
}

func byteToLocationDetails(body []byte) (LocationDetails, error) {
	var val LocationDetails
	err := json.Unmarshal(body, &val)
	if err != nil {
		return LocationDetails{}, errors.New(fmt.Sprintf("Encountered error while decoding Location Details: %v", err))
	}
	return val, nil
}
