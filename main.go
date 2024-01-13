package main

import (
	"fmt"
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

	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Type \"help\" for available commands")

	startRepl(&cfg)
}
