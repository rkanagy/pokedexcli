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

// GetNextLocationAreas returns the next set of locations
func GetNextLocationAreas(nextURL *string) (LocationAreas, error) {
	url, err := getNextURL(nextURL)
	if err != nil {
		return LocationAreas{}, err
	}

	locations, err := getLocationAreas(url)
	if err != nil {
		return LocationAreas{}, err
	}

	return locations, nil
}

// GetPreviousLocationAreas returns the previous set of locations
func GetPreviousLocationAreas(previousURL *string) (LocationAreas, error) {
	url, err := getPreviousURL(previousURL)
	if err != nil {
		return LocationAreas{}, err
	}

	locations, err := getLocationAreas(url)
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

func getLocationAreas(url string) (LocationAreas, error) {
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

	return locations, nil
}
