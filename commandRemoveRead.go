package main

import (
	"fmt"
)

func commandRemoveRead(c *config, args ...string) error {
	fmt.Println("User wants to remove a book from the read list")
	return nil
}