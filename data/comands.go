package data

import "fmt"

type cliCommand struct {
	Name        string
	Description string
	callback    func() error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	fmt.Println("Available commands:")
	return nil
}

func commandExit() error {
	fmt.Println("Available exit:")
	return nil
}
