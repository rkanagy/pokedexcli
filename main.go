package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/rkanagy/pokedexcli/internal/pokemon"
)

type cliCommand struct {
	name        string
	description string
	callback    func(parameters ...string) error
}

type pokedexType map[string]pokemon.Pokemon

var pokemonAPI pokemon.API = pokemon.NewAPI()
var pokedex pokedexType = make(pokedexType, 10)

func main() {
	commands := initializeCliCommands()

	// The Read-Eval-Print loop for the CLI
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")

	for reader.Scan() {
		text := cleanInput(reader.Text())
		splitText := strings.Split(text, " ")
		commandName := splitText[0]
		arguments := splitText[1:]

		// interpret commands
		if command, exists := commands[commandName]; exists {
			err := command.callback(arguments...)
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
		"explore": {
			name:        "explore",
			description: "Display the encountered Pokemon found at given location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Captures a Pokemon based on higher experience points making it more difficult",
			callback:    commandCatch,
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

func commandHelp(parameters ...string) error {
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

func commandExit(parameters ...string) error {
	os.Exit(0)
	return nil
}

func commandMap(parameters ...string) error {
	locations, err := pokemonAPI.GetLocationAreas(pokemon.Next)
	if err != nil {
		return err
	}

	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandMapb(parameters ...string) error {
	locations, err := pokemonAPI.GetLocationAreas(pokemon.Previous)
	if err != nil {
		return err
	}

	for _, result := range locations.Results {
		fmt.Println(result.Name)
	}

	return nil
}

func commandExplore(parameters ...string) error {
	if len(parameters) == 0 {
		return errors.New("No location area name was entered")
	}

	locationArea := parameters[0]
	location, err := pokemonAPI.GetLocationArea(locationArea)
	if err != nil {
		return err
	}

	fmt.Println("Exploring " + locationArea + "...")
	fmt.Println("Found Pokemon:")
	for _, pokemonEncounter := range location.PokemonEncounters {
		fmt.Println(" - " + pokemonEncounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(parameters ...string) error {
	if len(parameters) <= 0 {
		return errors.New("No Pokemon name was entered")
	}

	name := parameters[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := pokemonAPI.Capture(name)
	if err != nil {
		return err
	}
	if pokemon == nil {
		fmt.Printf("%s escaped!\n", name)
	} else {
		fmt.Printf("%s was caught!\n", name)
		pokedex[name] = *pokemon
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
