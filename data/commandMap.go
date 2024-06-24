package data

import (
	"errors"
	"fmt"
)

func CommandMapf(cfg *Config) error {
	res, err := cfg.PokeapiClient.GetLocationList(cfg.NextLocationsURL)

	if err != nil {
		return err
	}

	cfg.NextLocationsURL = res.Next
	cfg.PrevLocationsURL = res.Previous

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func CommandMapb(cfg *Config) error {

	if cfg.PrevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.PokeapiClient.GetLocationList(cfg.PrevLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locationResp.Next
	cfg.PrevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
