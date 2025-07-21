package main

import (
	"fmt"
)

func main() {
	text := "Hello, World!"
	fmt.Println("input: " + text)
	for _, word := range cleanInput(text) {
		fmt.Println(word)
	}
}
