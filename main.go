package main

import (
	"fmt"
	"time"

	"github.com/hatrnuhn/pokedexcli/internal/config"
	"github.com/hatrnuhn/pokedexcli/internal/pokeapi"
	"github.com/hatrnuhn/pokedexcli/internal/utils"
)

func main() {
	cfg := config.Config{
		PokeapiClient: pokeapi.NewClient(time.Hour / 2),
	}

	allPokemonsName, _ := utils.GetAllPokemons(&cfg)

	fmt.Println("Welcome to Pokedex!")
	fmt.Println("Type \"help\" for available commands")

	startRepl(&cfg, allPokemonsName)
}
