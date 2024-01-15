package main

import (
	"fmt"
)

func commandExplore(cfg *config, areaName string, _ map[string]bool) error {
	if areaName == "" {
		fmt.Println("Usage: explore area_name")
		return nil
	}

	resp, err := cfg.pokeapiClient.ListPokemonsInArea(areaName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n", areaName)
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
