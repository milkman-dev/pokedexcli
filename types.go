package main

import "github.com/milkman-dev/pokedexcli/pokeapi"

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	client   pokeapi.Client
	location *string
	next     *string
	previous *string
}
