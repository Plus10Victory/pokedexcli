package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Hello, World!"
	fmt.Println("input: " + text)
	for _, word := range cleanInput(text) {
		fmt.Println(word)
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
