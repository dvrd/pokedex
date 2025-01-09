package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		shouldContinue := scanner.Scan()
		// Retrieve and print the current line
		rawCommand := scanner.Text()
		trimmedCommand := strings.TrimSpace(rawCommand)
		loweredCommand := strings.ToLower(trimmedCommand)
		commandParts := strings.Fields(loweredCommand)
		command := commandParts[0]
		fmt.Println("Your command was:", command)

		// Check for errors during scanning
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
