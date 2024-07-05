package pokeapi

import (
	"net/http"
	"time"

	"github.com/Divyue30597/pokedex/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		cache: *pokecache.NewCache(60 * time.Second),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
