package main

import (
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
