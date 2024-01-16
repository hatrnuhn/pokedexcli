package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/hatrnuhn/pokedexcli/internal/config"
)

func commandCatch(cfg *config.Config, pokemon string, allPokemons map[string]bool) error {
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

	resp, err := cfg.PokeapiClient.PokemonReq(pokemon)
	if err != nil {
		return err
	}

	// checks whether user input exists in user's Pokedex
	_, ok := cfg.UserPokedex.Get(pokemon)
	if ok {
		return fmt.Errorf("%v is already in your Pokedex", pokemon)
	}

	fmt.Println("Throwing a Pokeball at ", pokemon)

	// chance gets harder diminishingly as base exp increases
	prob := 0.65 * math.Exp(-0.0112*float64(resp.BaseExperience))

	rFloat := r.Float64()
	isSuccess := rFloat < prob

	if isSuccess {
		fmt.Println(pokemon, "was caught!")

		// add pokemon to user's Pokedex when caught
		err = cfg.UserPokedex.Add(pokemon, resp)
		if err != nil {
			return err
		}
	} else {
		fmt.Println(pokemon, " escaped!")
	}

	return nil
}
