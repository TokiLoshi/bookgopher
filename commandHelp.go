package main

import (
	"fmt"
)

func commandHelp(c *config, args ...string) error {
	fmt.Println("User wants help")
	return nil
}