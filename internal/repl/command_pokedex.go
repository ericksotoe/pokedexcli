package repl

import "fmt"

func commandPokedex(cfg *Config, args []string) error {
	poke := cfg.Pokedex

	fmt.Println("Your Pokedex:")
	for _, pokemon := range poke {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
