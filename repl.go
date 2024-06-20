package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hw/go-repl/data"
)

func startRepl() {

	var exit = false
	for !exit {

		rd := bufio.NewReader(os.Stdin)
		fmt.Print("pokedex > ")
		message, err := rd.ReadString('\n')
		messageWithoutLineBreak := strings.TrimSpace(message)
		if err != nil {
			fmt.Println(err)
		}

		scanner := bufio.NewScanner(strings.NewReader(messageWithoutLineBreak))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
		res := data.GetCommands()
		command, exists := res[messageWithoutLineBreak]
		if exists {
			command.Callback()
		} else {
			fmt.Println("can't find the command ")
		}
		fmt.Print("\n")
	}
}
