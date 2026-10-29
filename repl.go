package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"

	"github.com/nitai-b/pokedexcli/internal/pokeapi"
	"github.com/nitai-b/pokedexcli/internal/pokecache"
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
	callback    func(cfg *cmdConfig, arguments []string) error
}

type cmdConfig struct {
	commands map[string]cliCommand
	Next     string
	Previous string
	cache    *pokecache.Cache
	pokemon  map[string]pokeapi.Pokemon
}

func commandExit(cfg *cmdConfig, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *cmdConfig, args []string) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
		`)
	return nil
}

func commandMap(cfg *cmdConfig, args []string) error {
	url := cfg.Next
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	}

	mapLocation, err := pokeapi.GetLocationData(url, cfg.cache)
	if err != nil {
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

func commandMapBack(cfg *cmdConfig, args []string) error {
	if cfg.Previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}
	url := cfg.Previous

	mapLocation, err := pokeapi.GetLocationData(url, cfg.cache)
	if err != nil {
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

func commandExplore(cfg *cmdConfig, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("location name arguments missing")
	}

	fmt.Printf("\nExploring %s...\n", args[0])

	url := "https://pokeapi.co/api/v2/location-area/" + args[0]

	locationDetails, err := pokeapi.GetLocationDetails(url, cfg.cache)
	if err != nil {
		return err
	}

	fmt.Printf("Found Pokemon:\n")
	for _, location := range locationDetails.PokemonEncounters {
		fmt.Printf("\t- %s\n", location.Pokemon.Name)
	}

	return nil
}

func commandCatch(cfg *cmdConfig, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("pokemon name argument missing")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])

	pokemon, err := pokeapi.GetPokemonDetails(args[0], cfg.cache)
	if err != nil {
		return err
	}

	const catchThreshold = 40

	caught := rand.IntN(pokemon.BaseExperience+1) < catchThreshold
	if caught {
		cfg.pokemon[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemon.Name)
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func commandInspect(cfg *cmdConfig, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("pokemon name argument is missing")
	}

	if pokemon, ok := cfg.pokemon[args[0]]; !ok {
		return fmt.Errorf("you have not caught this pokemon")
	} else {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, stat := range pokemon.Stats {
			fmt.Printf("\t- %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, pokeType := range pokemon.Types {
			fmt.Printf("\t- %s\n", pokeType.Type.Name)
		}

	}

	return nil
}
