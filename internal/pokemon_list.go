package internal

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonList(location string) (ResPokemonListFromLocation, error) {
	// url := BASE_URL + "location-area/" + location
	url := BASE_URL + "location-area/" + location

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResPokemonListFromLocation{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResPokemonListFromLocation{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ResPokemonListFromLocation{}, err
	}

	defer res.Body.Close()

	data := ResPokemonListFromLocation{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return ResPokemonListFromLocation{}, err
	}

	return data, nil

}
