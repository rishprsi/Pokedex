package main

import (
	"fmt"

	"www.github.com/rishprsi/PokedexAPI"
)

func commandMap(state *config) error {
	if state.Next == "" {
		state.Next = "https://pokeapi.co/api/v2/location-area/"
	}

	var areaStruct PokedexAPI.LocationAreaList
	areaStruct = state.client.ListLocationAreas()
	for _, area := range areaStruct.LocationList {
		fmt.Println(area.name)
	}
	state.Next = areaStruct.Next
	state.Prev = areaStruct.Prev

	return nil
}

func commandMapb(state *config) error {
	if state.Prev == "" {
		fmt.Println("You're on the first page")
	}
	var areaStruct PokedexAPI.LocationAreaList
	areaStruct = state.client.ListLocationAreas()
	for _, area := range areaStruct.LocationList {
		fmt.Println(area.name)
	}
	state.Next = areaStruct.Next
	state.Prev = areaStruct.Prev

	return nil
}
