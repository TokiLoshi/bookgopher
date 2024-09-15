package books

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type BooksService interface {
	ConvertBooks(path string) ([]BooksRead, error)
}

type BooksHandler struct{}

func NewBooksHandler() *BooksHandler {
	return &BooksHandler{}
}

func (b *BooksHandler) ConvertBooks(path string) ([]BooksRead, error) {

	var booksRead []BooksRead
	if len(path) == 0 {
		return booksRead, fmt.Errorf("need a file path")
	}

	// Read in contents of csv 
	file, err := os.Open(path)
	if err != nil {
		return booksRead, fmt.Errorf("issue opening csv: %w", err)
	}
	defer file.Close()

	// Create new reader with CSV
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return booksRead, fmt.Errorf("error reading records: %w", err)
	}
	
	// Add records to struct 
	for item, record := range records {
		// skipping the header 
		if item == 0 {
			continue
		}
		exclusiveShelf := record[18] 
		// Parse in read books to the BooksRead Struct
		if exclusiveShelf == "read" {
			book := BooksRead {
				BookId: atoi(record[0]),
				Title: record[1],
				Author: record[2],
				AuthorLF: record[3],
				AdditionalAuthors: record[4],
				ISBN: cleanString(record[5]),
				ISBN13: cleanString(record[6]),
				MyRating: atof(record[7]),
				AverageRating: atof(record[8]),
				Publisher: record[9],
				Binding: record[10],
				PageNumbers: atoi(record[11]),
				YearPublished: record[12],
				OriginalPublicationYear: record[13],
				DateRead: parseTime(record[15]),
				MyReview: record[19],
			}
			booksRead = append(booksRead, book)
		}

		// ToDo: 
		// do the same for the ToRead struct
	
	}
	
	// Marshal the json 
	booksReadJSON, err := json.MarshalIndent(booksRead, "", "  ")
	if err != nil {
		return booksRead, fmt.Errorf("error marshalling the json")
	}
	
	// Save the json file to disk
	outfilePath := "./output/books_read.json"
	err = os.WriteFile(outfilePath, booksReadJSON, 0644)
	
	if err != nil {
		return booksRead, fmt.Errorf("error writing the json")

	}
	fmt.Println("Json saved successfully")
	// save to the outputs folder 
	return booksRead, nil
}

func atoi(s string) int {
	convertedInt, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("error converting string to int: %v", err)
		return 0
	}
	return convertedInt
}

func atof(s string) float64 {
	convertedFloat, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("error converting string to float64, %v", err)
		return 0.0
	}
	return convertedFloat
}

func cleanString(s string) string {
	cleanedString := strings.Trim(s, "=\"")
	return cleanedString
}

func parseTime(s string) time.Time  {
	dateRead, err := time.Parse("2006/01/02", s)
	if err != nil {
		fmt.Printf("error converting time: %v", err)
		return time.Time{}
	}
	return dateRead
}