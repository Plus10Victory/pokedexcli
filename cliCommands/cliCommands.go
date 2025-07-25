package cliCommands

type cliCommand struct {
	name        string
	description string
	Callback    func(*Config) error
}

type Config struct {
	Next     *string
	Previous *string
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    HelpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    ExitCommand,
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
	}
}
