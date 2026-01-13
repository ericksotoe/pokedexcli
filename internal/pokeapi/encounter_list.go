package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PokemonNames struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetEncounterList(locationURL string) (PokemonNames, error) {
	url := baseURL + "/location-area/" + locationURL

	val, ok := c.PokeCache.Get(url)
	if ok {
		pokeNames := PokemonNames{}
		err := json.Unmarshal(val, &pokeNames)
		if err != nil {
			return PokemonNames{}, err
		}
		return pokeNames, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonNames{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonNames{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return PokemonNames{}, fmt.Errorf("Response failed with status code: %d and \nbody:%s\n", res.StatusCode, body)
	}
	if err != nil {
		return PokemonNames{}, err
	}

	c.PokeCache.Add(url, body)
	pokeNames := PokemonNames{}
	err = json.Unmarshal(body, &pokeNames)
	if err != nil {
		return PokemonNames{}, err
	}
	return pokeNames, nil

}
