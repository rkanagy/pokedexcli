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
const locationAreaBaseURL string = "https://pokeapi.co/api/v2/location-area/"

// NamedAPIResource contains the name and the url to retrieve further information that resource
type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// LocationAreas Structures ---------------------------------------------------

// LocationAreas contains the fields returned from location-area endpoint
type LocationAreas struct {
	Count    int         `json:"count"`
	Next     *string     `json:"next"`
	Previous *string     `json:"previous"`
	Results  []ResultsNR `json:"results"`
}

// ResultsNR contains the list of names and urls of each location area
type ResultsNR struct {
	NamedAPIResource
}

// ----------------------------------------------------------------------------

// LocationArea Structures ----------------------------------------------------

// LocationArea contains the information for a single Location Area
type LocationArea struct {
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	GameIndex            int                   `json:"game_index"`
	ID                   int                   `json:"id"`
	Location             LocationNR            `json:"location"`
	Name                 string                `json:"name"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

// EncounterMethodNR is a Named Resource for EncounterMethod
type EncounterMethodNR struct {
	NamedAPIResource
}

// VersionNR is a Named Resource for Version
type VersionNR struct {
	NamedAPIResource
}

// EncounterVersionDetails contains details on the rate and version of an encounter
type EncounterVersionDetails struct {
	Rate    int       `json:"rate"`
	Version VersionNR `json:"version"`
}

// EncounterMethodRate contains details on the method and rate of an encounter
type EncounterMethodRate struct {
	EncounterMethod EncounterMethodNR         `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}

// LocationNR is a Named Resource for Location
type LocationNR struct {
	NamedAPIResource
}

// LanguageNR is a Named Resource for Language
type LanguageNR struct {
	NamedAPIResource
}

// Name contains details on a name and its language information
type Name struct {
	Language LanguageNR `json:"language"`
	Name     string     `json:"name"`
}

// PokemonNR is a Named Resource for Pokemon
type PokemonNR struct {
	NamedAPIResource
}

// EncounterConditionValueNR is a Named Resource for EncounterConditionValue
type EncounterConditionValueNR struct {
	NamedAPIResource
}

// Encounter contains details on an encounter
type Encounter struct {
	Chance          int                         `json:"chance"`
	ConditionValues []EncounterConditionValueNR `json:"condition_values"`
	MaxLevel        int                         `json:"max_level"`
	Method          EncounterMethodNR           `json:"method"`
	MinLevel        int                         `json:"min_level"`
}

// VersionEncounterDetail contains details on the version of a set of encounters
type VersionEncounterDetail struct {
	EncounterDetails []Encounter `json:"encounter_details"`
	MaxChance        int         `json:"max_chance"`
	Version          VersionNR   `json:"version"`
}

// PokemonEncounter contains details on a Pokemon encounter
type PokemonEncounter struct {
	Pokemon        PokemonNR                `json:"pokemon"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

// ----------------------------------------------------------------------------

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

// GetLocationAreas returns the location areas corresponding to
// the direction (Previous or Next) passed in.
func (p *API) GetLocationAreas(direction int) (LocationAreas, error) {
	url, err := p.getURL(direction)
	if err != nil {
		return LocationAreas{}, err
	}
	body, err := p.httpGet(url)
	if err != nil {
		return LocationAreas{}, err
	}

	locations := LocationAreas{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return LocationAreas{}, err
	}
	p.updateConfig(locations)

	return locations, nil
}

// GetLocationArea returns the location area information for the given locationArea name
func (p *API) GetLocationArea(locationArea string) (LocationArea, error) {
	url := locationAreaBaseURL + locationArea
	body, err := p.httpGet(url)
	if err != nil {
		return LocationArea{}, err
	}

	location := LocationArea{}
	err = json.Unmarshal(body, &location)
	if err != nil {
		return LocationArea{}, err
	}

	return location, nil
}

func (p *API) getURL(direction int) (string, error) {
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

func (p *API) httpGet(url string) ([]byte, error) {
	// is the url in the cache?  If so, then get it from the cache,
	//otherwise do an HTTP Get on the url
	body, found := p.cache.Get(url)
	if !found {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		body, err = io.ReadAll(resp.Body)
		resp.Body.Close()
		if resp.StatusCode > 299 {
			msg := fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
			return nil, errors.New(msg)
		}
		if err != nil {
			return nil, err
		}
		p.cache.Add(url, body)
	}

	return body, nil
}

func (p *API) updateConfig(locations LocationAreas) {
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
