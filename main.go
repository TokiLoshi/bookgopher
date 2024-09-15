package main

import (
	"fmt"
	"time"

	"github.com/TokiLoshi/bookgopher/internal/gobookapi"
)

func main() {
	bookClient := gobookapi.NewClient(5 * time.Minute)
	cfg := &config {
		goapiClient: bookClient,
	}
	fmt.Println("Hello Gopher")
	startRepl(cfg)
}