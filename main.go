package main

import (
	"bufio"
	"fmt"
	"github.com/dvrd/pokedex/internal/commands"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		shouldContinue := scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		command, found := commands.Get()[input[0]]

		if found {
			err := command.Callback()
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

func cleanInput(text string) []string {
	var input []string

	trimmedInput := strings.TrimSpace(text)
	for _, word := range strings.Fields(trimmedInput) {
		input = append(input, strings.ToLower(word))
	}

	return input
}
