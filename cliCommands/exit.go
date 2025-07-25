package cliCommands

import (
	"fmt"
	"os"
)

func ExitCommand(config *Config) error {
	_ = config
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
