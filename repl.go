package main

import (
	"fmt"

	"github.com/hw/go-repl/data"
	"github.com/hw/go-repl/pkg"
)

func startRepl() {
	var s = ""
	var exit = false
	for !exit {

		fmt.Print("pokedex > ")
		fmt.Scanln(&s)
		fmt.Print("\n")
		res := data.GetCommands()
		command := res[s]

		switch {
		case s == pkg.Help:
			command.Callback()

		case s == pkg.Exit:
			command.Callback()
			exit = true

		default:
			fmt.Println("can't find the command")
		}
		fmt.Print("\n")
	}
}
