package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/dvrd/pokedex/internal/commands"
	"net/http"
	"os"
	"strings"
)

type Area struct {
	Url  string
	Name string
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

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
			description: "Displays name of next 20 location",
			callback:    commandMap,
		},
		"mapb": {
			name:        "map",
			description: "Displays name of previous 20 location",
			callback:    commandMapPrevious,
		},
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		shouldContinue := scanner.Scan()
		input := cleanInput(scanner.Text())
		command, found := getCommands()[input[0]]

		if found {
			err := command.callback()
			if err != nil {
				fmt.Println("ERROR: ", err)
			}
		} else {
			fmt.Println("Unknown command")
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
		}

		if !shouldContinue {
			break
		}
	}
}

var NextLocation string = "https://pokeapi.co/api/v2/location-area"
var PrevLocation string = ""

func getArea(url string) ([]Area, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("ERROR: request failed with status %v", res.Status)
	}

	type AreaResponse struct {
		Count    int
		Results  []Area
		Next     string
		Previous string
	}

	var jsonResponse AreaResponse
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&jsonResponse)
	if err != nil {
		return nil, err
	}

	if jsonResponse.Previous != "" {
		PrevLocation = jsonResponse.Previous
	} else if PrevLocation == "" {
		PrevLocation = NextLocation
	}

	if jsonResponse.Next != "" {
		NextLocation = jsonResponse.Next
	}

	return jsonResponse.Results, nil
}

func commandMap() error {
	locations, err := getArea(NextLocation)
	if err != nil {
		return err
	}

	for _, area := range locations {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapPrevious() error {
	if PrevLocation == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	locations, err := getArea(PrevLocation)
	if err != nil {
		return err
	}

	for _, area := range locations {
		fmt.Println(area.Name)
	}

	return nil
}

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()

	return nil
}

func cleanInput(text string) []string {
	var input []string

	trimmedInput := strings.TrimSpace(text)
	for _, word := range strings.Fields(trimmedInput) {
		input = append(input, strings.ToLower(word))
	}

	return input
}
