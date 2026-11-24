package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
	supportedCommands map[string]cliCommand
	next              *string
	previous          *string
}

func cleanInput(text string) []string {
	cleanedInput := make([]string, 0)
	splitWords := strings.SplitSeq(text, " ")

	for word := range splitWords {
		if strings.Trim(word, " ") != "" {
			cleanedInput = append(cleanedInput, strings.Trim(strings.ToLower(word), " "))
		}
	}

	return cleanedInput
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {

			userInput := cleanInput(scanner.Text())

			if val, ok := cfg.supportedCommands[userInput[0]]; ok {
				additionalArg := ""

				if len(userInput) > 1 {
					additionalArg = userInput[1]
				}

				err := val.callback(cfg, additionalArg)

				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}

		}

		if scanner.Err() != nil {
			_ = fmt.Errorf("Error occured while handling user input.")
		}
	}
}
