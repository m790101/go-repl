package internal

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationList(pageUrl *string) (RestApiLocation, error) {

	url := BASE_URL + "location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RestApiLocation{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RestApiLocation{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RestApiLocation{}, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return RestApiLocation{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return RestApiLocation{}, err
	}
	apiStruct := RestApiLocation{}
	err = json.Unmarshal(body, &apiStruct)
	if err != nil {
		return RestApiLocation{}, err
	}
	c.cache.Add(url, body)
	return apiStruct, nil

}
