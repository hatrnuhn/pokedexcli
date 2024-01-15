package utils

import (
	"log"

	"github.com/hatrnuhn/pokedexcli/internal/config"
)

func GetAllPokemons(cfg *config.Config) (map[string]bool, error) {
	resp, err := cfg.PokeapiClient.PokemonsListReq()
	if err != nil {
		log.Fatal(err)
	}

	pokemonsList := make(map[string]bool)
	for _, pokemon := range resp.Results {
		pokemonsList[pokemon.Name] = true
	}

	return pokemonsList, nil
}
