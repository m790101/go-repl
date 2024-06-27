package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hw/go-repl/data"
)

func startRepl(cfg *data.Config) {

	var exit = false
	for !exit {

		rd := bufio.NewReader(os.Stdin)
		fmt.Print("pokedex > ")
		message, err := rd.ReadString('\n')
		messageWithoutLineBreak := strings.TrimSpace(message)
		if err != nil {
			fmt.Println(err)
		}
		inputSlice := []string{}
		scanner := bufio.NewScanner(strings.NewReader(messageWithoutLineBreak))
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			line := scanner.Text()
			inputSlice = append(inputSlice, line)
		}
		res := data.GetCommands()
		command, exists := res[inputSlice[0]]
		cfg.InputSlice = inputSlice
		if exists {
			command.Callback(cfg)
		} else {
			fmt.Println("can't find the command ")
		}
		fmt.Print("\n")
	}
}
