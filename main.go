package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/nitai-b/pokedexcli/internal/pokeapi"
	"github.com/nitai-b/pokedexcli/internal/pokecache"
)

func main() {
	cfg := cmdConfig{
		cache:   pokecache.NewCache(5 * time.Second),
		pokemon: map[string]pokeapi.Pokemon{},
		commands: map[string]cliCommand{
			"exit": {
				name:        "exit",
				description: "Exit the Pokedex",
				callback:    commandExit,
			},
			"help": {
				name:        "help",
				description: "Help and usage of Pokedex",
				callback:    commandHelp,
			},
			"map": {
				name:        "map",
				description: "Get list of locations on map",
				callback:    commandMap,
			},
			"mapb": {
				name:        "mapb",
				description: "Navigate to previous list of locations",
				callback:    commandMapBack,
			},
			"explore": {
				name:        "explore",
				description: "Explore a location",
				callback:    commandExplore,
			},
			"catch": {
				name:        "catch",
				description: "Catch a Pokemon",
				callback:    commandCatch,
			},
			"inspect": {
				name:        "inspect",
				description: "Inspect Pokemon you caught",
				callback:    commandInspect,
			},
			"pokedex": {
				name:        "pokedex",
				description: "See the Pokemon you caught",
				callback:    commandPokedex,
			},
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		command, exists := cfg.commands[words[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(&cfg, words[1:]); err != nil {
			fmt.Println(err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
	}
}
