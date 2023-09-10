package pokemon

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const defaultNextURL string = "https://pokeapi.co/api/v2/location-area"

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
type Config struct {
	nextURL     *string
	previousURL *string
}

const (
	// Previous is the direction for previous set of location areas
	Previous = iota

	// Next is the direction for next set of location areas
	Next = iota
)

// GetLocationAreas returns the location areas corresponding to
// the direction (Previous or Next) passed in.
func GetLocationAreas(config *Config, direction int) (LocationAreas, error) {
	url, err := getURL(config, direction)
	if err != nil {
		return LocationAreas{}, err
	}

	resp, err := http.Get(url)
	if err != nil {
		return LocationAreas{}, err
	}

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode > 299 {
		msg := fmt.Sprintf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
		return LocationAreas{}, errors.New(msg)
	}
	if err != nil {
		return LocationAreas{}, err
	}

	locations := LocationAreas{}
	err = json.Unmarshal(body, &locations)
	if err != nil {
		return LocationAreas{}, err
	}
	updateConfig(locations, config)

	return locations, nil
}

// GetNextLocationAreas returns the next set of locations
func getNextLocationAreas(config *Config) (LocationAreas, error) {
	locations, err := GetLocationAreas(config, Next)
	if err != nil {
		return LocationAreas{}, err
	}

	return locations, nil
}

// GetPreviousLocationAreas returns the previous set of locations
func getPreviousLocationAreas(config *Config) (LocationAreas, error) {
	locations, err := GetLocationAreas(config, Previous)
	if err != nil {
		return LocationAreas{}, err
	}

	return locations, nil
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

func getURL(config *Config, direction int) (string, error) {
	var err error
	var url string

	switch direction {
	case Previous:
		url, err = getPreviousURL(config.previousURL)
		if err != nil {
			return "", err
		}
	case Next:
		url, err = getNextURL(config.nextURL)
		if err != nil {
			return "", err
		}
	default:
		return "", errors.New("getLocationAreas: invalid direction argument")
	}

	return url, nil
}

func updateConfig(locations LocationAreas, config *Config) {
	config.nextURL = locations.Next
	config.previousURL = locations.Previous
}