package postgresql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/zarasfara/go-rest-template/internal/config"
)

func Connect(cfg *config.DBConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.Username, cfg.Password, cfg.Database, cfg.Port, cfg.SSLMode)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
