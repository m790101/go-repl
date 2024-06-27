package main

import (
	"time"

	"github.com/hw/go-repl/data"
	"github.com/hw/go-repl/internal"
)

func main() {
	pokeClient := internal.NewClient(5*time.Second, time.Minute*5)
	cfg := &data.Config{
		PokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
