package data

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CommandMap() error {
	res, err := http.Get(LOCATION_URL)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	apiStruct := ApiStruct{}
	errUnmarshal := json.Unmarshal(body, &apiStruct)
	if errUnmarshal != nil {
		fmt.Println(err)
	}

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	for _, location := range apiStruct.Results {
		fmt.Println(location.Name)
	}
	return nil
}
