package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	// create a new scanner that reads from std input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex>")

	for scanner.Scan() {
		line := scanner.Text()

		switch strings.ToLower(line) {
		case "help":
			commands["help"].callback()
		case "exit":
			commands["exit"].callback()
		default:
			fmt.Println("command not identifiable")
		}
	}
}

func commandHelp() error {
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	return nil
}

func commandExit() error {
	fmt.Println("Exiting Pokedex CLI")
	os.Exit(0)
	return nil
}
