package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/rkanagy/pokedexcli/internal/pokemon"
)

type cliCommand struct {
	name        string
	description string
	//callback    func(config *pokemon.Config) error
	callback func() error
}

var pokemonAPI pokemon.Pokemon = pokemon.NewPokemon()

func main() {
	commands := initializeCliCommands()
	//currentConfig := pokemon.Config{}

	// The Read-Eval-Print loop for the CLI
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")

	for reader.Scan() {
		text := cleanInput(reader.Text())

		// interpret commands
		if command, exists := commands[text]; exists {
			//err := command.callback(&currentConfig)
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
		"map": {
			name:        "map",
			description: "Display the next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 location areas",
			callback:    commandMapb,
		},
	}
}

func cleanInput(text string) string {
	output := strings.TrimSpace(text)
	output = strings.ToLower(output)
	return output
}

func errorHandler(err error) {
	fmt.Fprintln(os.Stderr, err)
}

func commandNotRecognized() {
	fmt.Fprintln(os.Stderr, "command not recognized")
}

// func commandHelp(config *pokemon.Config) error {
func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	commands := initializeCliCommands()
	sortedKeys := sortKeys(commands)
	for _, key := range sortedKeys {
		command := commands[key]
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	fmt.Println()

	return nil
}

// func commandExit(config *pokemon.Config) error {
func commandExit() error {
	os.Exit(0)
	return nil
}

// func commandMap(config *pokemon.Config) error {
func commandMap() error {
	//locations, err := pokemonAPI.GetLocationAreas(config, pokemon.Next)
	locations, err := pokemonAPI.GetLocationAreas(pokemon.Next)
	if err != nil {
		return err
	}

	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}

	return nil
}

// func commandMapb(config *pokemon.Config) error {
func commandMapb() error {
	//locations, err := pokemonAPI.GetLocationAreas(config, pokemon.Previous)
	locations, err := pokemonAPI.GetLocationAreas(pokemon.Previous)
	if err != nil {
		return err
	}

	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func sortKeys(mapToSort map[string]cliCommand) []string {
	keys := make([]string, 0, len(mapToSort))

	for key := range mapToSort {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	return keys
}
