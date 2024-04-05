package book

import (
	"github.com/jmoiron/sqlx"
	"github.com/zarasfara/go-rest-template/internal/entity"
)

type mysqlRepository struct {
	db *sqlx.DB
}

func NewMySqlRepository(db *sqlx.DB) Repository {
	return &postgresRepository{
		db: db,
	}
}

func (r mysqlRepository) GetAll() ([]entity.Book, error) {
	var books []entity.Book = []entity.Book{
		{
			Id: 1,
		},
	}

	return books, nil
}
