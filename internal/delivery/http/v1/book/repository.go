package book

import "github.com/zarasfara/go-rest-template/internal/entity"

type Repository interface {
	// GetAll returns collection of books
	GetAll() ([]entity.Book, error)
}
