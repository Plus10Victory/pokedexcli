package cliCommands

import (
	"github.com/Plus10Victory/pokedexcli/internal/pokedex"
)

type cliCommand struct {
	name        string
	description string
	Callback    func(*pokedex.Config) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    HelpCommand,
		},
		"map": {
			name:        "map",
			description: "Displays a list of 20 map locations",
			Callback:    MapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous list of 20 map locations",
			Callback:    MapbCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    ExitCommand,
		},
	}
}
