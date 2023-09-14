package pokemon

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

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
