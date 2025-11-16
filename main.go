package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		userInputArray := cleanInput(userInput)
		if len(userInputArray) > 0 {
			fmt.Printf("Your command was: %s\n", userInputArray[0])
		} else {
			fmt.Println("No command found")
		}
	}
}
