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
				errorHandler(err)
			}
		} else {
			commandNotRecognized()
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

func errorHandler(err error) {
	fmt.Fprintln(os.Stderr, err)
}

func commandNotRecognized() {
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "command not recognized")
	fmt.Fprintln(os.Stderr)
}

func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}
