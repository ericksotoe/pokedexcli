package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// holds information on location areas found in pokemon
type PokeMap struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationAreas(pageURL *string) (PokeMap, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeMap{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokeMap{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return PokeMap{}, fmt.Errorf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return PokeMap{}, err
	}

	pokemonLocations := PokeMap{}
	err = json.Unmarshal(body, &pokemonLocations)
	if err != nil {
		return PokeMap{}, err
	}
	return pokemonLocations, nil
}
