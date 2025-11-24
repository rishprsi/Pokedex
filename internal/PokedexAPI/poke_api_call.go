package pokedexapi

import (
	// "context"
	"errors"
	"fmt"
	"io"
)

func (c *Client) PokeApiCall(url string) ([]byte, error) {
	res, err := c.http.Get(url)
	if err != nil {
		errStr := fmt.Sprintf("ListLocationAreas failed with: %v", err)
		return []byte{}, errors.New(errStr)
	}

	result, err := io.ReadAll(res.Body)
	if err != nil {
		errStr := fmt.Sprintf("ListLocationAreas failed with: %v", err)
		return []byte{}, errors.New(errStr)
	}

	return result, nil
}
