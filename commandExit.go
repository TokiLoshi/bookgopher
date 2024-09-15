package main

import (
	"fmt"
	"os"
)

func commandExit(c *config, args ...string) error {
	fmt.Println("Byeeeee")
	os.Exit(0)
	return nil
}