package books

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/TokiLoshi/bookgopher/internal/utils"
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
				BookId: utils.Atoi(record[0]),
				Title: record[1],
				Author: record[2],
				AuthorLF: record[3],
				AdditionalAuthors: record[4],
				ISBN: utils.CleanString(record[5]),
				ISBN13: utils.CleanString(record[6]),
				MyRating: utils.Atof(record[7]),
				AverageRating: utils.Atof(record[8]),
				Publisher: record[9],
				Binding: record[10],
				PageNumbers: utils.Atoi(record[11]),
				YearPublished: record[12],
				OriginalPublicationYear: record[13],
				DateRead: utils.ParseTime(record[15]),
				MyReview: record[19],
			}
			booksRead = append(booksRead, book)
		} else if exclusiveShelf == "to-read" {
			fmt.Printf("To Read with title: %v and record 16: %v\n", record[1], record[16])
			book := ToRead {
				BookId: utils.Atoi(record[0]),
				Title: record[1],
				Author: record[2],
				AuthorLF: record[3],
				AdditionalAuthors: record[4],
				ISBN: utils.CleanString(record[5]),
				ISBN13: utils.CleanString(record[6]),
				AverageRating: utils.Atof(record[8]),
				Publisher: record[9],
				PageNumbers: utils.Atoi(record[11]),
				YearPublished: record[12],
				OriginalPublicationYear: record[13],
				DateAdded: utils.ParseTime(record[15]),
			}
			toBeRead = append(toBeRead, book)
		} else {
			continue
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



