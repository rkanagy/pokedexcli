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
	commands := initializeCliCommands()

	// The Read-Eval-Print loop for the CLI
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")

	for reader.Scan() {
		text := cleanInput(reader.Text())

		// interpret commands
		if command, exists := commands[text]; exists {
			err := command.callback()
			if err != nil {
				// handle error
				fmt.Fprintln(os.Stderr, err)
			}
		} else {
			// invalid command
			fmt.Fprintln(os.Stderr, "command not recognized")
		}

		fmt.Print("Pokedex > ")
	}
}

func initializeCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for name, command := range initializeCliCommands() {
		fmt.Printf("%s: %s\n", name, command.description)
	}

	fmt.Println()

	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}
