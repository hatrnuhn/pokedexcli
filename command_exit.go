package main

import (
	"os"

	"github.com/hatrnuhn/pokedexcli/internal/config"
)

func commandExit(cfg *config.Config, _ string, _ map[string]bool) error {
	os.Exit(0)
	return nil
}
