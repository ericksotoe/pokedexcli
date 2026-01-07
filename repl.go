package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan() // advances the scanner
		userInput := scanner.Text()
		userInputSlice := cleanInput(userInput) 
		fmt.Printf("Your command was: %s\n", userInputSlice[0])
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	res := strings.Fields(text)
	return res
}