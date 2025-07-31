package pokedex

import (
	"github.com/Plus10Victory/pokedexcli/internal/pokeapi"
)

type Config struct {
	PokeapiClient        pokeapi.Client
	NextLocationsURL     *string
	PreviousLocationsURL *string
}
