package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

func cleanInput(text string) []string {
	output := strings.Split(strings.ToLower(strings.Trim(text, " ")), " ")
	return output
}

func pokeballCatch(base_experience int) bool {
	n1, n2 := rand.IntN(base_experience), int(base_experience/2)
	if n1 > n2 {
		return true
	}

	return false
}

func inspectPokedex(c *config, pokemon string) error {
	if p, ok := c.catchedPokemon[pokemon]; ok {
		fmt.Println("Name: ", p.Name)
		fmt.Println("Height: ", p.Height)
		fmt.Println("Weight: ", p.Weight)
	} else {
		fmt.Println("You have not caught that pokemon yet...")
	}

	return nil
}
