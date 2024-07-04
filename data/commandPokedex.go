package data

import "fmt"

func CommandPokedex(cfg *Config) error {
	list := cfg.CatchList.GetAll()
	if len((list)) == 0 {
		fmt.Println("No pokemon caught yet")
		return nil
	}
	for _, p := range list {
		fmt.Println(" - " + p)
	}
	return nil
}
