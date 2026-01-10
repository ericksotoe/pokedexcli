package pokeapi

import (
	"net/http"
	"time"

	"github.com/ericksotoe/pokedexcli/internal/pokecache"
)

type Client struct {
	httpClient http.Client
	PokeCache  *pokecache.Cache
}

func NewClient(timeout time.Duration, cache *pokecache.Cache) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		PokeCache: cache,
	}
}
