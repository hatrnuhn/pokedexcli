package main

import "fmt"

func commandHelp(cfg *config, _ string, _ map[string]bool) error {
	defer fmt.Printf("==================================================\n\n")
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Usage:")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf("Pokedex>%s - %s\n", cmd.name, cmd.description)
	}

	return nil
}
