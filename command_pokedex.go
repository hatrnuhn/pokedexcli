package main

import (
	"fmt"

	"github.com/hatrnuhn/pokedexcli/internal/config"
)

func commandPokedex(cfg *config.Config, _ string, _ map[string]bool) error {
	userPokedex, err := cfg.UserPokedex.GetUserPokedex()
	if err != nil {
		return err
	}

	fmt.Println("Your Pokedex:")

	for _, val := range userPokedex {
		fmt.Printf("  - %v\n", val.Name)
	}
	return nil
}
