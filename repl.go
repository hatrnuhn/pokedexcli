package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	// create a new scanner that reads from std input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex>")

		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)

		// if user doesn't type anything but presses enter, continue
		if len(cleaned) == 0 {
			continue
		}

		// take first word user input as command
		userCommand := cleaned[0]

		// get available commands
		availableCommands := getCommands()

		// print invalid command when command is unavailable then continue the input loop
		command, ok := availableCommands[userCommand]
		if !ok {
			fmt.Println("invalid command")
			continue
		}

		// else execute the function
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exits Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "lists next list of location location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "prints previous list of location areas",
			callback:    commandMapb,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
