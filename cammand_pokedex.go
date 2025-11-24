package main

import (
	"fmt"
)

func commandPokedex(state *config, args []string) error {
	fmt.Println("Your Pokedex:")
	if len(state.Pokedex) == 0 {
		fmt.Println("Is empty :((")
	} else {
		for _, pokemon := range state.Pokedex {
			fmt.Printf("  -- %s\n", pokemon.Name)
		}
	}
	return nil
}
