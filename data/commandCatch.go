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
	input := cfg.InputSlice[1]
	// api get info from name
	fmt.Printf("\n")
	fmt.Println("Catching " + input + "...")
	fmt.Printf("\n")

	res, err := cfg.PokeApiClient.GetPokemonInfo(input)
	if err != nil {
		fmt.Println("error when geting the info...")
		return err
	}
	// use base_experience to get the catch chance
	baseEx := res.BaseExperience

	// fmt.Printf("baseEx is %d\n", baseEx)
	t := time.Now().UnixNano()
	r1 := rand.New(rand.NewSource(t))
	num := r1.Intn(200)

	// determine whether its a success
	time.Sleep(1 * time.Second)
	if num > baseEx {
		fmt.Println(res.Name + " was caught!")
		cfg.CatchList.Add(res)
		// catchList := cfg.CatchList.GetAll()
		// fmt.Println(res.Stats[0].BaseStat)
		// fmt.Println(res.Stats[0].Stat.Name)
		// fmt.Println(catchList)

	} else {
		fmt.Println("pikachu escaped!")
	}
	// save to the memory

	return nil
}
