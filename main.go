package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		cmd := cleanInput(scanner.Text())
		if len(cmd[0]) == 0 {
			continue
		}
		if cmd[0] == "quit" {
			break
		}

		fmt.Println("Your command was:", cmd[0])
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("error:", err)
	}
}
