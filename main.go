package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/milkman-dev/pokedexcli/pokeapi"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &config{
		client:   pokeapi.NewClient(5 * time.Second),
		next:     nil,
		previous: nil,
	}

	for {
		fmt.Print("Pokedex > ")
		if ok := scanner.Scan(); !ok {
			break
		}

		command := scanner.Text()
		commands := getCommands()

		if c, ok := commands[command]; ok {
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
