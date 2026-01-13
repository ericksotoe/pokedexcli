package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pokemon struct {
	Name   string `json:"name"`
	BaseXP int    `json:"base_experience"`
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name + "/"

	val, ok := c.PokeCache.Get(url)
	if ok {
		poke := Pokemon{}
		err := json.Unmarshal(val, &poke)
		if err != nil {
			return Pokemon{}, err
		}
		return poke, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("Response failed with status code: %d and \nbody:%s\n", res.StatusCode, body)

	}
	if err != nil {
		return Pokemon{}, err
	}

	c.PokeCache.Add(url, body)
	poke := Pokemon{}
	err = json.Unmarshal(body, &poke)
	if err != nil {
		return Pokemon{}, err
	}
	return poke, nil
}
