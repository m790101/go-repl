package main

import (
	"time"

	"github.com/hw/go-repl/data"
	"github.com/hw/go-repl/internal"
)

func main() {
	pokeClient := internal.NewClient(5 * time.Second)
	cfg := &data.Config{
		PokeapiClient: pokeClient, // Fix: Change httpClient to HTTPClient
	}
	startRepl(cfg)
}
