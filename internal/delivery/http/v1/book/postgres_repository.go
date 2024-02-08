package book

import (
	"github.com/jmoiron/sqlx"
	"github.com/zarasfara/go-rest-template/internal/entity"
)

type postgresRepository struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) Repository {
	return &postgresRepository{
		db: db,
	}
}

func (r postgresRepository) GetAll() ([]entity.Book, error) {
	var books []entity.Book = []entity.Book {
		{
			Id: 1,
		},
	}

	return books, nil
}
