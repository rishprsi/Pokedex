package main

import (
	"fmt"
	"os"
)

func commandExit(state *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	defer os.Exit(0)
	return nil
}
