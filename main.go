package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hatrnuhn/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour / 2),
	}

	allPokemonsName, _ := getAllPokemons(&cfg)

	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Type \"help\" for available commands")

	startRepl(&cfg, allPokemonsName)
}

func getAllPokemons(cfg *config) (map[string]bool, error) {
	resp, err := cfg.pokeapiClient.PokemonsListReq()
	if err != nil {
		log.Fatal(err)
	}

	pokemonsList := make(map[string]bool)
	for _, pokemon := range resp.Results {
		pokemonsList[pokemon.Name] = true
	}

	return pokemonsList, nil
}
