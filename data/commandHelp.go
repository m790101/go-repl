package data

import "fmt"

func CommandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("usage:")
	fmt.Println("")
	for _, command := range GetCommands() {
		fmt.Println(command.Name + ": " + command.Description)
	}
	return nil
}
