package main

import (
	"fmt"
	"time"

	"github.com/TokiLoshi/bookgopher/internal/gobookapi"
	"github.com/TokiLoshi/bookgopher/internal/models/books"
)

func main() {
	bookClient := gobookapi.NewClient(5 * time.Minute)
	cfg := &config {
		goapiClient: bookClient,
		Books: books.NewBooksHandler(),
	}
	fmt.Println("Hello Gopher")
	startRepl(cfg)
}