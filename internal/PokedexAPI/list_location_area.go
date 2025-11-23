package pokedexapi

import (
	// "context"
	"errors"
	"fmt"
	"io"
)

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaList struct {
	Prev         string         `json:"previous"`
	Next         string         `json:"next"`
	LocationList []LocationArea `json:"results"`
}

func (c *Client) ListLocationAreas(url string) ([]byte, error) {
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
