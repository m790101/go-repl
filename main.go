package main

import (
	"fmt"

	"github.com/hw/go-repl/data"
	"github.com/hw/go-repl/pkg"
)

func main() {
	var s = ""
	var exit = false
	for !exit {

		fmt.Print("pokedex > ")
		fmt.Scanln(&s)
		fmt.Print("\n")
		res := data.GetCommands()
		switch {
		case s == pkg.Help:
			fmt.Println("Welcome to the Pokedex!")

			for _, command := range res {

				fmt.Println(command.Name + ": " + command.Description)
			}
		case s == pkg.Exit:
			exit = true
		}
		fmt.Print("\n")
	}

}
