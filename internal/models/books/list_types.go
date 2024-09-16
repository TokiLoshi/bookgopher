package books

import "time"

// Records structure in CSV
// 0 - BookId
// 1 - Title
// 2 - Author
// 3 - Author l-f
// 4 - Additional Authors
// 5 - ISBN
// 6 - ISBN13
// 7 - My Rating
// 8 - Average Rating
// 9 - Publisher
// 10 - Binding
// 11 - Number of Pages
// 12 - Year Published
// 13 - Original Publication Year
// 14 - Date Read
// 15 - Date Added
// 16 - Bookshelves
// 17 - Bookshelves with positions
// 18 - Exclusive Shelf
// 19 - My reivew
// 20 - Spoiler
// 21 - Private notes
// 22 - Read Count
// 23 - Owned Copies

type BooksRead struct {
	BookId                  int
	Title                   string
	Author                  string
	AuthorLF                string
	AdditionalAuthors       string
	ISBN                    string
	ISBN13                  string
	MyRating                float64
	AverageRating           float64
	Publisher               string
	Binding                 string
	PageNumbers             int
	YearPublished           string
	OriginalPublicationYear string
	DateRead                time.Time
	MyReview                string

}

type ToRead struct {
	BookId                  int
	Title                   string
	Author                  string
	AuthorLF                string
	AdditionalAuthors       string
	ISBN                    string
	ISBN13                  string
	AverageRating           float64
	Publisher               string
	PageNumbers             int
	YearPublished           string
	OriginalPublicationYear string
	DateAdded               time.Time
}