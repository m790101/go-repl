package data

import "fmt"

func CommandExplore(cfg *Config) error {
	if len(cfg.InputSlice) < 2 {
		fmt.Println("no location to explore")
		return nil
	}
	location := cfg.InputSlice[1]

	fmt.Printf("\n")
	fmt.Println("Exploring " + location + "...")
	fmt.Printf("\n")

	pokemonList, err := cfg.PokeapiClient.GetPokemonList(location)

	if err != nil {

		fmt.Println("error getting pokemon list")
		return err
	}

	for _, pokemon := range pokemonList.PokemonEncounters {
		fmt.Println("- " + pokemon.Pokemon.Name)
	}
	return nil
}

func lens(s []string) {
	panic("unimplemented")
}
