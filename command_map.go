package main

import (
	"encoding/json"
	"fmt"
)

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaList struct {
	Prev         string         `json:"previous"`
	Next         string         `json:"next"`
	LocationList []LocationArea `json:"results"`
}

func commandMap(state *config, args []string) error {
	if state.Next == "" {
		state.Next = state.BaseURL + "location-area/"
	}
	state.Page += 1

	body, err := cacheHit(state, state.Next)
	if err != nil {
		return err
	}

	areaStruct, err := byteToList(body)
	if err != nil {
		return err
	}
	fmt.Println("\nOn Page:", state.Page, "\n")
	for _, area := range areaStruct.LocationList {
		fmt.Println(area.Name)
	}
	state.Next = areaStruct.Next
	state.Prev = areaStruct.Prev

	return nil
}

func commandMapb(state *config, args []string) error {
	if state.Prev == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	state.Page -= 1

	body, err := cacheHit(state, state.Prev)
	if err != nil {
		return err
	}

	areaStruct, err := byteToList(body)
	if err != nil {
		return err
	}

	fmt.Println("\nOn Page:", state.Page, "\n")
	for _, area := range areaStruct.LocationList {
		fmt.Println(area.Name)
	}
	state.Next = areaStruct.Next
	state.Prev = areaStruct.Prev

	return nil
}

func byteToList(body []byte) (LocationAreaList, error) {
	var locations LocationAreaList
	err := json.Unmarshal(body, &locations)
	if err != nil {
		return LocationAreaList{}, err
	}
	return locations, nil
}
