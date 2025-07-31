package cliCommands

import (
	"fmt"

	"github.com/Plus10Victory/pokedexcli/internal/pokedex"
)

func HelpCommand(config *pokedex.Config) error {
	_ = config
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for key, value := range GetCommands() {
		fmt.Println(key + ": " + value.description)
	}
	fmt.Println()
	return nil
}
