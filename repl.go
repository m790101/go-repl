package main

import (
	"fmt"

	"github.com/hw/go-repl/data"
)

func startRepl() {
	var s = ""
	var exit = false
	for !exit {

		fmt.Print("pokedex > ")
		fmt.Scanln(&s)
		fmt.Print("\n")
		res := data.GetCommands()
		command, exists := res[s]
		if exists {
			command.Callback()
		} else {
			fmt.Println("can't find the command")
		}
		fmt.Print("\n")
	}
}
