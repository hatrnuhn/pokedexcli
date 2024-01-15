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

	fmt.Println("Throwing a Pokeball at ", pokemon)

	resp, err := cfg.PokeapiClient.PokemonReq(pokemon)
	if err != nil {
		return err
	}

	prob := 0.65 * math.Exp(-0.0112*float64(resp.BaseExperience))
	fmt.Println(prob)
	rFloat := r.Float64()
	fmt.Println(rFloat)
	isSuccess := rFloat < prob

	if isSuccess {
		fmt.Println(pokemon, "was caught!")
	} else {
		fmt.Println(pokemon, " escaped!")
	}

	return nil
}
