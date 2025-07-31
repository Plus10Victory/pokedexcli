package cliCommands

import (
	"fmt"
	"os"

	"github.com/Plus10Victory/pokedexcli/internal/pokedex"
)

func ExitCommand(config *pokedex.Config) error {
	_ = config
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
