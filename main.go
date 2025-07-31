package main

import (
	"time"

	"github.com/Plus10Victory/pokedexcli/internal/pokeapi"
	"github.com/Plus10Victory/pokedexcli/internal/pokedex"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	config := &pokedex.Config{
		PokeapiClient: pokeClient,
	}

	startRepl(config)
}
