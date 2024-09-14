package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name string 
	description string 
	callback func() error
}

func startRepl() {
	fmt.Println("===========================")
	fmt.Println("Starting GoPher Books REPL")
	fmt.Println("===========================")
	reader := bufio.NewScanner(os.Stdin)
	for  {
		fmt.Print(">>>> ")
		reader.Scan()

		// Get the command line args 
		
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := string(words[0])
		if command, ok := getCommands()[commandName]; ok {
			err := command.callback()
			if err != nil {
				fmt.Printf("Error from %v command\n", command)
			}
			continue
		} else {
		fmt.Println("Unknown command")
		continue
	}
	}
}

func cleanInput(text string) []string {
	cleanedInput := strings.ToLower(text)
	// Split the words by space
	words := strings.Fields(cleanedInput)
	return words
}

func getCommands() map[string] cliCommand {
	return map[string]cliCommand {
	"help" : {
		name: "help", 
		description: "displays a help message",
		callback: commandHelp,
	},
	"exit" : {
		name: "exit", 
		description: "exits the program",
		callback: commandExit,
	},
}
}