package pokedexapi

import (
	"net/http"
)

type Client struct {
	http *http.Client
}

func PokedexClient() *Client {
	httpClient := http.Client{}
	client := Client{
		http: &httpClient,
	}
	return &client
}
