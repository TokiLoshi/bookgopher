package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Byeeeee")
	os.Exit(0)
	return nil
}