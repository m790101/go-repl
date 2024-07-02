package data

import (
	"fmt"
	"math/rand"
	"time"
)

func CommandCatch(cfg *Config) error {
	if len(cfg.InputSlice) < 2 {
		fmt.Println("no location to explore")
		return nil
	}
	name := cfg.InputSlice[1]
	// api get info from name
	fmt.Printf("\n")
	fmt.Println("Catching " + name + "...")
	fmt.Printf("\n")

	res, err := cfg.PokeapiClient.GetPokemonInfo(name)
	if err != nil {
		fmt.Println("error when geting the info...")
	}
	// use base_experience to get the catch chance
	baseEx := res.BaseExperience

	fmt.Printf("baseEx is %d\n", baseEx)
	t := time.Now().UnixNano()
	r1 := rand.New(rand.NewSource(t))
	num := r1.Intn(200)
	// fmt.Println(num)
	// determine whether its a success
	if num > baseEx {
		fmt.Println("pikachu was caught!")
	} else {
		fmt.Println("pikachu escaped!")
	}
	// save to the memory

	return nil
}
