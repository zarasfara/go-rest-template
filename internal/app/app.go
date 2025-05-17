package app

import (
	"log/slog"

	"github.com/zarasfara/go-rest-template/internal/config"
	"github.com/zarasfara/go-rest-template/internal/delivery/http"
	"github.com/zarasfara/go-rest-template/internal/server"
	"github.com/zarasfara/go-rest-template/pkg/database/postgresql"
)

func Run(cfg *config.Config) {
	// connect to database
	db, err := postgresql.Connect(&cfg.DB)
	if err != nil {
		slog.Error("failed connect to database", "error", err)
	}

	// build http
	server := server.NewServer(cfg, http.NewRouter(db))

	// start http
	server.Run()
}
