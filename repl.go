package main

import "strings"

func cleanInput(text string) []string {
	words := strings.Fields(text)

	var final []string
	for _, w := range words {
		final = append(final, strings.ToLower(w))
	}

	return final
}
