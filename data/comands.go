package data

import (
	"github.com/hw/go-repl/internal"
	"github.com/hw/go-repl/internal/pokelist"
)

type Config struct {
	PokeApiClient    internal.Client
	NextLocationsURL *string
	PrevLocationsURL *string
	InputSlice       []string
	CatchList        pokelist.PokemonCatchList
}

type cliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    func(config *Config) error { return CommandExit() },
		},
		"map": {
			Name:        "map",
			Description: "Displays a map of the region",
			Callback:    CommandMapf,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Get the previous page of locations",
			Callback:    CommandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "explore the location",
			Callback:    CommandExplore,
		},
		"inspect": {
			Name:        "inspect",
			Description: "inspect a pokemon",
			Callback:    CommonInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "check your pokemons",
			Callback:    CommandPokedex,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch a pokemon",
			Callback:    CommandCatch,
		},
	}
}
