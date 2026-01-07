package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Cli command struct that stores functions depending on user input
type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// This function holds all the commands that can be entered into the cli
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "",
		},
	}
}

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan() // advances the scanner
		userInput := scanner.Text()
		userInputSlice := cleanInput(userInput)

		// if the user just presses enter reload the prompt
		if len(userInputSlice) == 0 {
			fmt.Print("Unknown command\n")
			continue
		}

		command, ok := getCommands()[userInputSlice[0]]
		if ok {
			if err := command.callback(); err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}
