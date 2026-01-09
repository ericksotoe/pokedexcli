package repl

import (
	"fmt"
)

func commandMap(cfg *Config, args []string) error {
	pokemonMap, err := cfg.PokeApiClient.GetLocationAreas(cfg.Next)
	if err != nil {
		return err
	}

	cfg.Next = pokemonMap.Next
	cfg.Previous = pokemonMap.Previous

	for _, cityData := range pokemonMap.Results {
		fmt.Println(cityData.Name)
	}
	return nil
}

func commandMapB(cfg *Config, args []string) error {
	if cfg.Previous == nil {
		return fmt.Errorf("you're on the first page")
	}

	pokemonMap, err := cfg.PokeApiClient.GetLocationAreas(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = pokemonMap.Next
	cfg.Previous = pokemonMap.Previous

	for _, cityData := range pokemonMap.Results {
		fmt.Println(cityData.Name)
	}

	return nil
}
