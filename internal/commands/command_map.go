package commands

import (
	"fmt"
	"pokedexapi"
)

func commandMap(state *config) error {
	if state.Next == "" {
		state.Next = "https://pokeapi.co/api/v2/location-area/"
	}

	var areaStruct pokedexapi.LocationAreaList
	areaStruct, err := cacheHit(state, state.Next)
	if err != nil {
		return err
	}
	for _, area := range areaStruct.LocationList {
		fmt.Println(area.Name)
	}
	state.Next = areaStruct.Next
	state.Prev = areaStruct.Prev

	return nil
}

func commandMapb(state *config) error {
	if state.Prev == "" {
		fmt.Println("You're on the first page")
		return nil
	}
	areaStruct, err := cacheHit(state, state.Prev)
	if err != nil {
		return err
	}
	for _, area := range areaStruct.LocationList {
		fmt.Println(area.Name)
	}
	state.Next = areaStruct.Next
	state.Prev = areaStruct.Prev

	return nil
}
