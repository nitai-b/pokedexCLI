package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/nitai-b/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	words := strings.Fields(text)

	for i, w := range words {
		words[i] = strings.ToLower(w)
	}

	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*cmdConfig) error
}

type cmdConfig struct {
	commands map[string]cliCommand
	Next     string
	Previous string
}

func commandExit(cfg *cmdConfig) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *cmdConfig) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
		`)
	return nil
}

func commandMap(cfg *cmdConfig) error {
	url := cfg.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("error closing response body: ", err)
		}
	}()

	var mapLocation pokeapi.LocationResponse
	if err := json.NewDecoder(resp.Body).Decode(&mapLocation); err != nil {
		return err
	}

	cfg.Next = mapLocation.Next
	cfg.Previous = ""
	if mapLocation.Previous != nil {
		cfg.Previous = *mapLocation.Previous
	}

	for _, location := range mapLocation.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapBack(cfg *cmdConfig) error {
	if cfg.Previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}
	url := cfg.Previous

	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Println("error closing response body: ", err)
		}
	}()

	var mapLocation pokeapi.LocationResponse
	if err := json.NewDecoder(resp.Body).Decode(&mapLocation); err != nil {
		return err
	}

	cfg.Next = mapLocation.Next
	if mapLocation.Previous != nil {
		cfg.Previous = *mapLocation.Previous
	}

	for _, location := range mapLocation.Results {
		fmt.Println(location.Name)
	}

	return nil
}
