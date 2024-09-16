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
	ConvertBooks(path string) ([]BooksRead, []ToRead, error)
}

type BooksHandler struct{}

func NewBooksHandler() *BooksHandler {
	return &BooksHandler{}
}

func (b *BooksHandler) ConvertBooks(path string) ([]BooksRead, []ToRead, error) {

	var booksRead []BooksRead
	var toBeRead []ToRead
	if len(path) == 0 {
		return booksRead, toBeRead, fmt.Errorf("need a file path")
	}

	// Read in contents of csv 
	file, err := os.Open(path)
	if err != nil {
		return booksRead, toBeRead, fmt.Errorf("issue opening csv: %w", err)
	}
	defer file.Close()

	// Create new reader with CSV
	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return booksRead, toBeRead, fmt.Errorf("error reading records: %w", err)
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
		} else {
			book := ToRead {
				BookId: atoi(record[0]),
				Title: record[1],
				Author: record[2],
				AuthorLF: record[3],
				AdditionalAuthors: record[4],
				ISBN: cleanString(record[5]),
				ISBN13: cleanString(record[6]),
				AverageRating: atof(record[8]),
				Publisher: record[9],
				PageNumbers: atoi(record[11]),
				YearPublished: record[12],
				OriginalPublicationYear: record[13],
				DateAdded: parseTime(record[16]),
			}
			toBeRead = append(toBeRead, book)
		}
	
	}
	
	// Marshal the json 
	booksReadJSON, err := json.MarshalIndent(booksRead, "", "  ")
	if err != nil {
		return booksRead, toBeRead, fmt.Errorf("error marshalling read list into json")
	}

	booksToReadJSON, err := json.MarshalIndent(toBeRead, "", " ")
	if err != nil {
		return booksRead, toBeRead, fmt.Errorf("error marshalling to be read list into json")
	}
	
	// Save the json file to disk
	readOutfilePath := "./output/books_read.json"
	err = os.WriteFile(readOutfilePath, booksReadJSON, 0644)
	
	if err != nil {
		return booksRead, toBeRead, fmt.Errorf("error writing the read list into json")

	}
	fmt.Println("Json saved successfully")

	toReadOutfilePath := "./output/books_to_read.json"
	err = os.WriteFile(toReadOutfilePath, booksToReadJSON, 0644)
	if err != nil {
		return booksRead, toBeRead, fmt.Errorf("error writing to be read list into json")
	}
	// save to the outputs folder 
	return booksRead, toBeRead, nil
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