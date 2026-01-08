package main

import(
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Brando-1212/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct{
	name	    string
	description string
	callback    func(*config) error
}



func repl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0{
			continue
		}

		commandword := words[0]
		
		command, exists := getCommands()[commandword]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string{

	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}





func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name: 		  "map",
			description:  "Get the next page locations",
			callback:     commandMapf,
		},
		"mapb": {
			name:         "mapb",
			description:  "Get the previous page of locations",
			callback:     commandMapb,
		},
		"exit": {
			name:         "exit",
			description:  "Exit the Pokedex",
			callback:     commandExit,
		},
	}
}
