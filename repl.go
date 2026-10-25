package main

import (
	"fmt"
	"os"
	"strings"
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
