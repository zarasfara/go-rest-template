package book

import (
	"encoding/json"
	"io"
	"os"

	"github.com/zarasfara/go-rest-template/internal/entity"
)

const file = "./books.json"

type fileRepository struct {
}

func NewFileRepository() Repository {
	return &fileRepository{}
}

// GetBooks implements Repository.
func (*fileRepository) GetBooks() ([]entity.Book, error) {
	f, err := os.OpenFile(file, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	var books []entity.Book
	if err := json.Unmarshal(data, &books); err != nil {
		return nil, err
	}

	return books, nil
}
