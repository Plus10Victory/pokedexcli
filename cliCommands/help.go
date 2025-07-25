package cliCommands

import (
	"fmt"
)

func HelpCommand(config *Config) error {
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
