package pokelist

type Pokemon struct {
	Name string
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

func (c *PokemonCatchList) Add(name string) {
	// c.mux.Lock()
	c.CatchList[name] = Pokemon{
		Name: name,
	}
	// defer c.mux.Unlock()
}

func (c *PokemonCatchList) GetAll() []string {
	slice := []string{}
	for _, p := range c.CatchList {
		slice = append(slice, p.Name)
	}
	return slice
}
