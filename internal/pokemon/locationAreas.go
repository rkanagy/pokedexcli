package pokemon

import (
	"encoding/json"
	"errors"
)

const defaultNextURL string = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"

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
