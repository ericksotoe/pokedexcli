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

func GetLocationAreas(url string) (PokeMap, error) {
	res, err := http.Get(url)
	if err != nil {
		return PokeMap{}, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return PokeMap{}, fmt.Errorf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		return PokeMap{}, err
	}

	pokemonMap := PokeMap{}
	err = json.Unmarshal(body, &pokemonMap)
	if err != nil {
		return PokeMap{}, err
	}
	return pokemonMap, nil
}
