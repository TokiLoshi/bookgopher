package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/TokiLoshi/bookgopher/internal/gobookapi"
	"github.com/TokiLoshi/bookgopher/internal/models/books"
)

type cliCommand struct {
	name string 
	description string 
	callback func(*config, ...string) error
}

type config struct {
	goapiClient gobookapi.Client 
	Books books.BooksService 
}

func startRepl(cfg *config) {
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
		args := []string{}
		if command, ok := getCommands()[commandName]; ok {
			err := command.callback(cfg, args...)
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
	"read" : {
		name: "markRead",
		description: "marks a book as read",
		callback: commandReaddit,
	},
	"new-to-read" : {
		name: "addToRead",
		description: "adds book to reading list",
		callback: commandAddToList,
	},
	"delete-read" : {
		name: "removeRead",
		description: "removes a book from read",
		callback: commandRemoveRead,
	},
	"delete-to-read" : {
		name: "removeList",
		description: "removes a book from list",
		callback: commandRemoveList,
	},
	"parse-csv" : {
		name: "parseCSV",
		description: "takes uploaded goodreads reading lists and to read list and parses it",
		callback: commandParseCSV,
	},
}
}