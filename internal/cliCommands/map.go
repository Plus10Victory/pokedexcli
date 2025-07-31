package cliCommands

import (
	"errors"
	"fmt"

	"github.com/Plus10Victory/pokedexcli/internal/pokedex"
)

func MapCommand(config *pokedex.Config) error {
	locationsResp, err := config.PokeapiClient.ListLocations(config.NextLocationsURL)
	if err != nil {
		return err
	}

	if locationsResp.Next != nil {
		config.NextLocationsURL = locationsResp.Next
	}
	if locationsResp.Previous != nil {
		config.PreviousLocationsURL = locationsResp.Previous
	}

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func MapbCommand(config *pokedex.Config) error {
	if config.PreviousLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := config.PokeapiClient.ListLocations(config.PreviousLocationsURL)
	if err != nil {
		return err
	}

	config.NextLocationsURL = locationsResp.Next
	config.PreviousLocationsURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}
