package main

import (
	"time"

	"github.com/hw/go-repl/data"
	"github.com/hw/go-repl/internal"
	"github.com/hw/go-repl/internal/pokelist"
)

func main() {
	pokeClient := internal.NewClient(5*time.Second, time.Minute*5)
	catchList := pokelist.NewCatchList()
	cfg := &data.Config{
		PokeApiClient: pokeClient,
		CatchList:     catchList,
	}
	startRepl(cfg)
}
