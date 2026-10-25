package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	commands := map[string]cliCommand{
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

		command, exists := commands[words[0]]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(); err != nil {
			fmt.Println(err)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
	}
}
