package book

import "github.com/zarasfara/go-rest-template/internal/entity"

type Repository interface {
	// GetBooks returns collection of books
	GetBooks() ([]entity.Book, error)
}
