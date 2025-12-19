package main

import "strings"

func cleanInput(text string) []string {
	output := strings.Split(strings.ToLower(strings.Trim(text, " ")), " ")
	return output
}
