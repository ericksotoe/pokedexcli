package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ericksotoe/pokedexcli/internal/pokeapi"
)

// this struct will hold the next and previous urls for the map command
type Config struct {
	Pokedex       map[string]pokeapi.Pokemon
	PokeApiClient pokeapi.Client
	Next          *string
	Previous      *string
}

// Cli command struct that stores functions depending on user input
type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config, args []string) error
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
			description: "Displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explores the pokemon found at the given location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch the named pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects the named pokemon and prints stats if caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all the pokemon you have caught",
			callback:    commandPokedex,
		},
	}
}

func StartRepl(cfg *Config) {
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

		cmdName := userInputSlice[0]
		args := userInputSlice[1:]

		command, ok := getCommands()[cmdName]
		if ok {
			if err := command.callback(cfg, args); err != nil {
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
