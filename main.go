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
		scanner.Scan() // advances the scanner
		userInput := scanner.Text()
		userInputSlice := cleanInput(userInput) 
		fmt.Printf("Your command was: <%s>", userInputSlice[0])
	}
}


