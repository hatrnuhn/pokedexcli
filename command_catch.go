package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, pokemon string, allPokemons map[string]bool) error {
	if pokemon == "" {
		fmt.Println("Usage: catch pokémon_name, replace spaces with \"-\" if any")
		return nil
	}

	if !allPokemons[pokemon] {
		fmt.Println("pokémon is not found, make sure to type it right!")
		return nil
	}

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	fmt.Println("Throwing a Pokeball at ", pokemon)

	resp, err := cfg.pokeapiClient.PokemonReq(pokemon)
	if err != nil {
		return err
	}

	probability := 1 / (1 + math.Log10(float64(resp.BaseExperience)+1))
	isSuccess := r.Float64() < probability

	if isSuccess {
		fmt.Println(pokemon, "was caught!")
	} else {
		fmt.Println(pokemon, " escaped!")
	}

	return nil
}
