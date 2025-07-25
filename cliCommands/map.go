package cliCommands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type mapResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []locationArea `json:"results"`
}

type locationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func MapCommand(config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if config.Next != nil {
		if strings.Contains(*config.Next, "/location-area/") {
			url = *config.Next
		}
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d\n", res.StatusCode)
	}

	var mapRes mapResponse
	err = json.Unmarshal(body, &mapRes)
	if err != nil {
		return fmt.Errorf("Error: unmarshalling failed: %w\n", err)
	}

	if mapRes.Previous != nil {
		config.Previous = mapRes.Previous
	}

	if mapRes.Next != nil {
		fmt.Println()
		fmt.Println("updating Next: " + *mapRes.Next)
		fmt.Println()
		config.Next = mapRes.Next
	}

	for _, location := range mapRes.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func MapbCommand(config *Config) error {
	url := "https://pokeapi.co/api/v2/location-area/"

	if config.Previous != nil {
		if strings.Contains(*config.Previous, "/location-area/") {
			url = *config.Previous
		} else {
			return fmt.Errorf("Error: Previous entry was not map, try 'map' instead\n")
		}
	} else {
		return fmt.Errorf("Error: No Previous entries saved, try 'map' or 'help' instead\n")
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d\n", res.StatusCode)
	}

	var mapRes mapResponse
	err = json.Unmarshal(body, &mapRes)
	if err != nil {
		return fmt.Errorf("Error: unmarshalling failed: %w\n", err)
	}

	if mapRes.Previous != nil {
		config.Previous = mapRes.Previous
	}

	if mapRes.Next != nil {
		config.Next = mapRes.Next
	}

	for _, location := range mapRes.Results {
		fmt.Println(location.Name)
	}

	return nil
}
