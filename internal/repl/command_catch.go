package repl

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *Config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Pass in a pokemon name after catch")
	}

	poke, err := cfg.PokeApiClient.GetPokemon(args[0])
	if err != nil {
		return fmt.Errorf("invalid pokemon name")
	}
	pokemonName := poke.Name
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	difficulty := chanceToCatch(poke.BaseXP)
	chance := rand.Intn(100)
	if chance < difficulty {
		fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.Pokedex[pokemonName] = poke
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
}

func chanceToCatch(baseXP int) (difficultyToCatch int) {
	switch {
	case baseXP <= 50:
		return 80
	case baseXP <= 100:
		return 60
	case baseXP <= 200:
		return 40
	case baseXP <= 250:
		return 25
	case baseXP <= 270:
		return 15
	default:
		return 5
	}
}
