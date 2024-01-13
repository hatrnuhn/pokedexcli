package main

import (
	"fmt"

	"github.com/hatrnuhn/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Type \"help\" for available commands")

	startRepl(&cfg)
}
