package main

import (
	"encoding/json"
	"fmt"

	"github.com/hatrnuhn/pokedexcli/internal/config"
	"github.com/hatrnuhn/pokedexcli/internal/pokeapi"
)

func commandInspect(cfg *config.Config, pokemon string, allPokemons map[string]bool) error {
	if pokemon == "" {
		fmt.Println("Usage: inspect pokémon_name, replace spaces with \"-\" if any")
		return nil
	}

	if !allPokemons[pokemon] {
		fmt.Println("pokémon is not found, make sure to type it right!")
		return nil
	}

	key := pokeapi.BaseURL + "/pokemon/" + pokemon

	val, exists := cfg.PokeapiClient.Cache.Get(key)

	if !exists {
		return fmt.Errorf("you don't have %v in your pokedex", pokemon)
	}

	p := pokeapi.PokemonResp{}

	err := json.Unmarshal(val, &p)
	if err != nil {
		return err
	}

	str := fmt.Sprintf("Name: %v\nHeight: %v\nWeight: %v\nStats:\n",
		p.Name, p.Height, p.Weight,
	)

	for _, stat := range p.Stats {
		str += fmt.Sprintf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	str += "Types:\n"

	for _, typ := range p.Types {
		str += fmt.Sprintf("  -%v\n", typ.Type.Name)
	}

	fmt.Println(str)

	return nil
}
