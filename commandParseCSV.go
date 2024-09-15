package main

import (
	"fmt"
)

// Followed this tutorial https://cihanozhan.medium.com/converting-csv-data-to-json-with-go-b6f683644b38
func commandParseCSV(cfg *config, args ...string) error {
	fmt.Println("Let's parse the csv")
	// get the path to the csv 
	goodReads := "goodreadsExport.csv"
	bookPath := "./input/" + goodReads
	books, err := cfg.Books.ConvertBooks(bookPath)
	if err != nil {
		fmt.Println("Ok we're back from books")
	}
	fmt.Println(books)
	// parse the input file 
	// handle the errors
	// parse into json 
	// save file 
	return nil
}