package repl

import (
	"fmt"

	"github.com/ericksotoe/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config, args []string) error {
	pokemonMap, err := pokeapi.GetLocationAreas(cfg.Next)
	if err != nil {
		return err
	}

	if pokemonMap.Previous != nil {
		cfg.Previous = *pokemonMap.Previous
	} else {
		cfg.Previous = ""
	}

	if pokemonMap.Next != nil {
		cfg.Next = *pokemonMap.Next
	}

	for _, cityData := range pokemonMap.Results {
		fmt.Println(cityData.Name)
	}

	return nil
}
