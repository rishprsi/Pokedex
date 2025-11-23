module Pokedex

require internal/pokedexapi v0.0.0

require internal/pokecache v0.0.0

replace internal/pokedexapi => ./internal/PokedexAPI

replace internal/pokecache => ./internal/pokecache

go 1.25.2
