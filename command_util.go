package main

func cacheHit(state *config, url string) ([]byte, error) {
	val, ok := state.cache.Get(url)
	if !ok {
		apiVal, err := state.client.PokeApiCall(url)
		if err != nil {
			return []byte{}, err
		}
		state.cache.Add(url, apiVal)
		return apiVal, nil

	} else {
		return val, nil
	}
}
