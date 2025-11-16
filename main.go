package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := initCommands()
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		userInputArray := cleanInput(userInput)
		if len(userInputArray) > 0 {
			if command, ok := commands[userInputArray[0]]; ok {
				err := command.callback()
				if err != nil {
					fmt.Println("Invalid command usage")
				}
			} else {
				fmt.Println("Command not found, use 'help' for a list of available commands")
			}
		} else {
			fmt.Println("No valid command found use help for getting help")
		}
	}
}
