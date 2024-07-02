package internal

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(name string) (RestPokemonInfo, error) {
	// url := BASE_URL + "location-area/" + location
	url := BASE_URL + "pokemon/" + name

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RestPokemonInfo{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RestPokemonInfo{}, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return RestPokemonInfo{}, err
	}

	defer res.Body.Close()

	data := RestPokemonInfo{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return RestPokemonInfo{}, err
	}

	return data, nil

}
