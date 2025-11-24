package commands

import "testing"

func TestCommand(t *testing.T) {
	commandList := []string{"help", "exit", "map"}
	commands := initCommands()
	for _, commandName := range commandList {
		if _, ok := commands[commandName]; !ok {
			t.Errorf("The command %s should be available but is not currently available", commandName)
			t.Fail()
		}
	}
}
