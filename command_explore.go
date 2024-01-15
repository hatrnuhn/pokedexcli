package main

import (
	"fmt"

	"github.com/hatrnuhn/pokedexcli/internal/config"
)

func commandExplore(cfg *config.Config, areaName string, _ map[string]bool) error {
	if areaName == "" {
		fmt.Println("Usage: explore area_name")
		return nil
	}

	resp, err := cfg.PokeapiClient.ListPokemonsInArea(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n", areaName)
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
