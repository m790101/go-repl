package data

import "fmt"

func CommonInspect(cfg *Config) error {
	fmt.Println("Inspecting...")
	fmt.Println("")
	// check catch list if its empty
	if len(cfg.CatchList.GetAll()) == 0 {
		fmt.Println("No pokemon to inspect")
	} else {
		// get information of the pokemon
		input := cfg.InputSlice[1]
		p := cfg.CatchList.Get(input)
		fmt.Println("Name: " + p.Name)
		fmt.Println("Height: ", p.Height)
		fmt.Println("Weight: ", p.Weight)
		fmt.Println("")
		fmt.Println("Stats: ")
		for _, s := range p.Stats {
			fmt.Printf(" -%v:  %d\n", s.Stat.Name, s.BaseStat)
			// fmt.Println(s.BaseStat)
		}
		fmt.Println("")
		fmt.Println("Types: ")
		for _, t := range p.Types {
			fmt.Println(" -" + t.Type.Name)
		}
	}
	return nil
}
