package main

import (
	"fmt"
	"os"
)

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Show map areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Show previous map areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore area pokemons",
			callback:    commandExplore,
		},
	}

	return commands
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp(c *config) error {
	commands := getCommands()
	fmt.Println("Usage: ")
	fmt.Println("")
	for _, cmd := range commands {
		fmt.Printf("%s: %s", cmd.name, cmd.description)
		fmt.Println("")
	}

	return nil
}

func commandMap(c *config) error {
	locations, err := c.client.GetLocations(c.next)
	if err != nil {
		return err
	}

	c.next, c.previous = locations.Next, locations.Previous
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(c *config) error {
	locations, err := c.client.GetLocations(c.previous)
	if err != nil {
		return err
	}

	c.next, c.previous = locations.Next, locations.Previous
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExplore(c *config) error {
	pokemons, err := c.client.GetPokemons(c.location)
	if err != nil {
		return err
	}

	fmt.Println("Exploring location...")
	for _, v := range pokemons.PokemonEncounters {
		pokemon := v.Pokemon
		fmt.Printf("- %s", pokemon.Name)
		fmt.Println()
	}

	return nil
}
