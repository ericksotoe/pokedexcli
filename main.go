package main

import (
	"time"

	"github.com/ericksotoe/pokedexcli/internal/pokeapi"
	"github.com/ericksotoe/pokedexcli/internal/pokecache"
	"github.com/ericksotoe/pokedexcli/internal/repl"
)

func main() {
	cache := pokecache.NewCache(time.Second * 3)
	pokeClient := pokeapi.NewClient(5*time.Second, cache)

	cfg := repl.Config{
		PokeApiClient: pokeClient,
	}
	repl.StartRepl(&cfg)
}
