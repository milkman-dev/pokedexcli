package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/milkman-dev/pokedexcli/pokeapi"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{
		client:   pokeapi.NewClient(5*time.Second, 10*time.Second),
		location: nil,
		next:     nil,
		previous: nil,
	}

	for {
		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); !ok {
			break
		}

		input := scanner.Text()
		command := strings.Split(input, " ")
		commands := getCommands()

		if len(command) == 2 {
			switch command[0] {
			case "explore":
				location := command[1]
				cfg.location = &location
			case "catch":
				pokemon := command[1]
				cfg.pokemon = &pokemon
			case "inspect":
				pokemon := command[1]
				cfg.pokemon = &pokemon
			}
		} else {
			switch command[0] {
			case "explore":
				fmt.Println("Must specify location")
				continue
			case "catch":
				fmt.Println("Must specify pokemon")
				continue
			}
		}

		if c, ok := commands[command[0]]; ok {
			err := c.callback(cfg)
			if err != nil {
				fmt.Println("Error calling command", err)
			}
		} else {
			fmt.Printf("Unknow command: %s", command)
			fmt.Println()
		}
	}
}
