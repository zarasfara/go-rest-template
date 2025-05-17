package postgresql

import (
	"fmt"
	"net/url"

	"github.com/jmoiron/sqlx"
	"github.com/zarasfara/go-rest-template/internal/config"
)

func Connect(cfg *config.DBConfig) (*sqlx.DB, error) {
	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.Username, cfg.Password),
		Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	}
	q := dsn.Query()
	q.Add("sslmode", cfg.SSLMode)
	dsn.RawQuery = q.Encode()

	db, err := sqlx.Open("postgres", dsn.String())
	if err != nil {
		return nil, err
	}
	return db, nil
}
