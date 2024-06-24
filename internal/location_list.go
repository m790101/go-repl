package internal

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c Client) GetLocationList(pageUrl *string) (RestApiLocation, error) {
	url := BASE_URL + "location-area"
	if pageUrl != nil {
		url = *pageUrl
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
	return apiStruct, nil

}
