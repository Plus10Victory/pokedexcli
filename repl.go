package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Plus10Victory/pokedexcli/internal/cliCommands"
	"github.com/Plus10Victory/pokedexcli/internal/pokedex"
)

func startRepl(config *pokedex.Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := cliCommands.GetCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		} else {
			err := command.Callback(config)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
