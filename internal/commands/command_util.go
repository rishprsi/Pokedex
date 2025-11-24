package commands

import (
	"encoding/json"
	"pokedexapi"
)

func cacheHit(state *config, url string) (pokedexapi.LocationAreaList, error) {
	val, ok := state.cache.Get(url)
	if !ok {
		apiVal, err := state.client.ListLocationAreas(url)
		if err != nil {
			return pokedexapi.LocationAreaList{}, err
		}
		state.cache.Add(url, apiVal)
		return byteToList(apiVal)

	} else {
		return byteToList(val)
	}
}

func byteToList(body []byte) (pokedexapi.LocationAreaList, error) {
	var locations pokedexapi.LocationAreaList
	err := json.Unmarshal(body, &locations)
	if err != nil {
		return pokedexapi.LocationAreaList{}, err
	}
	return locations, nil
}
