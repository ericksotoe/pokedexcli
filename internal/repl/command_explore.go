package repl

import (
	"fmt"
)

func commandExplore(cfg *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Pass in a city name after explore")
	}
	pokeNames, err := cfg.PokeApiClient.GetEncounterList(args[0])
	if err != nil {
		return fmt.Errorf("invalid pokemon location area name")
	}
	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon:")
	for _, val := range pokeNames.PokemonEncounters {
		fmt.Printf(" - %s\n", val.Pokemon.Name)
	}
	return nil
}
