package books

import (
	"fmt"
)

type BooksService interface {
	ConvertBooks(path string) ([]BooksRead, error)
}

type BooksHandler struct{}

func NewBooksHandler() *BooksHandler {
	return &BooksHandler{}
}

func (b *BooksHandler) ConvertBooks(path string) ([]BooksRead, error) {
	fmt.Printf("Trying to convert from %v\n", path)
	return []BooksRead{}, nil
}