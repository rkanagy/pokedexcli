package pokemon

import "encoding/json"

const locationAreaBaseURL string = "https://pokeapi.co/api/v2/location-area/"

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
