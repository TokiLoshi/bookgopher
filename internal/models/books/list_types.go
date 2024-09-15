package books

import "time"

type BooksRead struct {
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
	DateRead                time.Time
	MyReview                string
	Genres                   []string
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