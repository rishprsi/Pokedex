package pokedexapi

import (
	// "context"
	"encoding/json"
	"errors"
	"fmt"
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

func (c *Client) ListLocationAreas(url string) (LocationAreaList, error) {
	res, err := c.http.Get(url)
	if err != nil {
		errStr := fmt.Sprintf("ListLocationAreas failed with: %v", err)
		return LocationAreaList{}, errors.New(errStr)
	}

	decoder := json.NewDecoder(res.Body)
	var result LocationAreaList
	err = decoder.Decode(&result)
	if err != nil {
		errStr := fmt.Sprintf("ListLocationAreas failed with: %v", err)
		return LocationAreaList{}, errors.New(errStr)
	}

	return result, nil
}
