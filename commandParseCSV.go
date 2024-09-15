package main

import (
	"fmt"
)

// Followed this tutorial https://cihanozhan.medium.com/converting-csv-data-to-json-with-go-b6f683644b38
func commandParseCSV(cfg *config, args ...string) error {
	
	
	goodReads := "goodreadsExport.csv"
	bookPath := "./input/" + goodReads
	fmt.Printf("Parsing CSV located at %v\n", bookPath)
	_, err := cfg.Books.ConvertBooks(bookPath)
	if err != nil {
		fmt.Println("something went wrong after trying to convert the books")
	}

	fmt.Println("CSV parsed successfully")
	return nil
}