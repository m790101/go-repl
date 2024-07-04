package pokelist

import "github.com/hw/go-repl/internal"

type Pokemon struct {
	Name   string
	Height int
	Stats  []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

type PokemonCatchList struct {
	CatchList map[string]Pokemon
}

func NewCatchList() PokemonCatchList {

	c := PokemonCatchList{
		CatchList: make(map[string]Pokemon),
	}

	return c
}

func (c *PokemonCatchList) Add(data internal.RestPokemonInfo) {

	c.CatchList[data.Name] = Pokemon{
		Name:   data.Name,
		Height: data.Height,
		Stats:  data.Stats,
		Types:  data.Types,
		Weight: data.Weight,
	}
}

func (c *PokemonCatchList) GetAll() []string {
	slice := []string{}
	for _, p := range c.CatchList {
		slice = append(slice, p.Name)
	}
	return slice
}

func (c *PokemonCatchList) Get(name string) Pokemon {

	return c.CatchList[name]
}
