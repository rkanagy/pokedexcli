package pokemon

import (
	"time"

	"github.com/rkanagy/pokedexcli/internal/pokecache"
)

// API contains cached responses from the Pokemon API
type API struct {
	cache  pokecache.Cache
	config config
}

// NewAPI creates a new Pokemon struct
func NewAPI() API {
	return API{
		cache:  pokecache.NewCache(5 * time.Minute),
		config: config{},
	}
}
