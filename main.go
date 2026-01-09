package main

import (
	"time"

	"github.com/ericksotoe/pokedexcli/internal/pokeapi"
	"github.com/ericksotoe/pokedexcli/internal/repl"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := repl.Config{
		PokeApiClient: pokeClient,
	}
	repl.StartRepl(&cfg)
}
