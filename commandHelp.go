package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("=================================")
	fmt.Println("            MENU                 ")
	fmt.Println("=================================")
	options := getCommands()
	count := 1
	for _, command := range options {
		fmt.Printf("%v. %v : %v\n", count, command.name, command.description)
		count++
	}
	return nil
}