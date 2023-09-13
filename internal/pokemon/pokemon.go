package pokemon

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rkanagy/pokedexcli/internal/pokecache"
)

const defaultNextURL string = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"

// LocationAreas contains the fields returned from location-area endpoint
type LocationAreas struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Results `json:"results"`
}

// Results contains the list of names and urls of each location area
type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Config contains pointers to the next and previous URLs
type config struct {
	nextURL     *string
	previousURL *string
}

const (
	// Previous is the direction for previous set of location areas
	Previous = iota

	// Next is the direction for next set of location areas
	Next = iota
)

// Pokemon contains cached responses from the Pokemon API
type Pokemon struct {
	cache  pokecache.Cache
	config config
}

// NewPokemon creates a new Pokemon struct
func NewPokemon() Pokemon {
	return Pokemon{
		cache:  pokecache.NewCache(5 * time.Minute),
		config: config{},
	}
}

// GetLocationAreas returns the location areas corresponding to
// the direction (Previous or Next) passed in.
func (p *Pokemon) GetLocationAreas(direction int) (LocationAreas, error) {
	url, err := p.getURL(direction)
	if err != nil {
		return LocationAreas{}, err
	}

	// is the url in the cache?  If so, then get it from the cache,
	//otherwise do an HTTP Get on the url
	body, found := p.cache.Get(url)
	if !found {
		resp, err := http.Get(url)
		if err != nil {
			return LocationAreas{}, err
		}

		body, err = io.ReadAll(resp.Body)
		resp.Body.Close()
		if resp.StatusCode > 299 {
			msg := fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
			return LocationAreas{}, errors.New(msg)
		}
		if err != nil {
			return LocationAreas{}, err
		}
		p.cache.Add(url, body)
	}

	locations := LocationAreas{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return LocationAreas{}, err
	}
	p.updateConfig(locations)

	return locations, nil
}

func (p *Pokemon) getURL(direction int) (string, error) {
	var err error
	var url string

	switch direction {
	case Previous:
		url, err = getPreviousURL(p.config.previousURL)
		if err != nil {
			return "", err
		}
	case Next:
		url, err = getNextURL(p.config.nextURL)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("getLocationAreas: invalid direction argument")
	}

	return url, nil
}

func (p *Pokemon) updateConfig(locations LocationAreas) {
	p.config.nextURL = locations.Next
	p.config.previousURL = locations.Previous
}

func getNextURL(nextURL *string) (string, error) {
	url := defaultNextURL
	if nextURL != nil {
		url = *nextURL
	}
	return url, nil
}

func getPreviousURL(previousURL *string) (string, error) {
	if previousURL == nil {
		return "", errors.New("At top of locations list")
	}

	return *previousURL, nil

}
