package book

import "github.com/zarasfara/go-rest-template/internal/entity"

type Repository interface {
	GetAll() ([]entity.Book, error)
}
