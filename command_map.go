package main

import (
	"errors"
	"fmt"

	"github.com/hatrnuhn/pokedexcli/internal/config"
)

func commandMap(cfg *config.Config, _ string, _ map[string]bool) error {
	resp, err := cfg.PokeapiClient.ListLocationAreas(cfg.NextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas: ")
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	cfg.NextLocationAreaURL = resp.Next
	cfg.PrevLocationAreaURL = resp.Previous

	return nil

}

func commandMapb(cfg *config.Config, _ string, _ map[string]bool) error {
	if cfg.PrevLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}
	resp, err := cfg.PokeapiClient.ListLocationAreas(cfg.PrevLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas: ")
	for _, area := range resp.Results {
		fmt.Printf("- %s\n", area.Name)
	}

	cfg.NextLocationAreaURL = resp.Next
	cfg.PrevLocationAreaURL = resp.Previous

	return nil

}
