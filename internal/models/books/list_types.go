package books

import "time"

// Records structure in CSV
// 1 - BookId
// 2 - Title
// 3 - Author
// 4 - Author l-f
// 5 - Additional Authors
// 6 - ISBN
// 7 - ISBN13
// 8 - My Rating
// 9 - Average Rating
// 10 - Publisher
// 11 - Binding
// 12 - Number of Pages
// 13 - Year Published
// 14 - Original Publication Year
// 15 - Date Read
// 16 - Date Added
// 17 - Bookshelves
// 18 - Bookshelves with positions
// 19 - Exclusive Shelf
// 20 - My reivew
// 21 - Spoiler
// 22 - Private notes
// 23 - Read Count
// 24 - Owned Copies

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
	AdditionalAuthors       []string
	ISBN                    string
	ISBN13                  string
	MyRating                float64
	AverageRating           float64
	Publisher               string
	Binding                 string
	PageNumbers             int
	YearPublished           string
	OriginalPublicationYear string
	DateAdded               time.Time
	Genres                  []string
}