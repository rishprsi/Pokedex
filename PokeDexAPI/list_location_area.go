package pokedexapi

import (
	// "context"
	"encoding/json"
	"errors"
	"fmt"
)

type LocationArea struct {
	name string
	url  string
}

type LocationAreaList struct {
	Prev         string
	Next         string
	LocationList []LocationArea
}

type PokedexLocations struct {
	count     int
	previous  string
	next      string
	locations []LocationArea
}

func (c *Client) ListLocationAreas(url string) (LocationAreaList, error) {
	res, err := c.http.Get(url)
	if err != nil {
		err_str := fmt.Sprintf("ListLocationAreas failed with: %v", err)
		return LocationAreaList{}, errors.New(err_str)
	}

	decoder := json.NewDecoder(res.Body)
	var result PokedexLocations
	err = decoder.Decode(&result)
	if err != nil {
		err_str := fmt.Sprintf("ListLocationAreas failed with: %v", err)
		return LocationAreaList{}, errors.New(err_str)
	}

	locationAreaList := LocationAreaList{
		Prev:         result.previous,
		Next:         result.next,
		LocationList: result.locations,
	}

	return locationAreaList, nil
}
