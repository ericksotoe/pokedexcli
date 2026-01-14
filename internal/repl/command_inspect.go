package repl

import "fmt"

func commandInspect(cfg *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Pass in a pokemon name after inspect")
	}
	pokemonName := args[0]
	pokemon, ok := cfg.Pokedex[pokemonName]
	if ok {
		fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\n", pokemonName, pokemon.Height, pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, poke_type := range pokemon.Types {
			fmt.Printf("  -%s\n", poke_type.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}
