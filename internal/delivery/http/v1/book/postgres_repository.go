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

func (r *postgresRepository) GetBooks() ([]entity.Book, error) {
	var books []entity.Book
	err := r.db.Select(&books, "SELECT * FROM books")
	if err!= nil {
		return nil, err
	}
	return books, nil
}
