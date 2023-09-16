package pokemon

import (
	"encoding/json"
	"math/rand"
	"time"
)

const pokemonBaseURL string = "https://pokeapi.co/api/v2/pokemon/"

// Capture captures a Pokemon based on base experience
func (p *API) Capture(name string) (*Pokemon, error) {
	pokemon, err := p.getPokemon(name)
	if err != nil {
		return nil, err
	}

	if p.isCaptured(pokemon) {
		return &pokemon, nil
	} else {
		return nil, nil
	}
}

func (p *API) getPokemon(name string) (Pokemon, error) {
	url := pokemonBaseURL + name
	body, err := p.httpGet(url)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}

func (p *API) isCaptured(pokemon Pokemon) bool {
	captureRate := p.calculateSuccessfulCaptureRate(pokemon)
	randomNumber := p.generateRandomNumber(100.0)

	if randomNumber <= captureRate {
		return true
	}
	return false
}

func (p *API) calculateSuccessfulCaptureRate(pokemon Pokemon) float64 {
	experience := pokemon.BaseExperience
	successfulCapturePercent := 100 - (float64(experience) / 2)
	if successfulCapturePercent < 5.0 {
		successfulCapturePercent = 5.0
	} else if successfulCapturePercent > 100 {
		successfulCapturePercent = 95.0
	}

	return successfulCapturePercent
}

func (p *API) generateRandomNumber(high float64) float64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Float64() * high

	return randomNumber
}
