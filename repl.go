package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Plus10Victory/pokedexcli/cliCommands"
)

func startRepl() {
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
			err := command.Callback()
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
