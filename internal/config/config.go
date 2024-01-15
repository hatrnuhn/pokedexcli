package config

import "github.com/hatrnuhn/pokedexcli/internal/pokeapi"

type Config struct {
	PokeapiClient       pokeapi.Client
	NextLocationAreaURL *string
	PrevLocationAreaURL *string
}
