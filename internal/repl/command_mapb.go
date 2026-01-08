package repl

import (
	"fmt"

	"github.com/ericksotoe/pokedexcli/internal/pokeapi"
)

func commandMapB(cfg *config, args []string) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	pokemonMap, err := pokeapi.GetLocationAreas(cfg.Previous)
	if err != nil {
		return err
	}
	if pokemonMap.Previous != nil {
		cfg.Previous = *pokemonMap.Previous
	} else {
		cfg.Previous = ""
		fmt.Println("you're on the first page")
	}

	if pokemonMap.Next != nil {
		cfg.Next = *pokemonMap.Next
	}

	for _, cityData := range pokemonMap.Results {
		fmt.Println(cityData.Name)
	}

	return nil
}
