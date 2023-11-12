package app

import (
	"github.com/sirupsen/logrus"
	"github.com/zarasfara/go-rest-template/internal/config"
	"github.com/zarasfara/go-rest-template/internal/delivery/http"
	"github.com/zarasfara/go-rest-template/internal/server"
	"github.com/zarasfara/go-rest-template/pkg/database/postgresql"
)

func Run(cfg *config.Config) {
	// connect to database
	db, err := postgresql.Connect(&cfg.DB)
	if err != nil {
		logrus.Fatalf("failed connect to database: %s", err.Error())
	}

	// build http
	serv := server.NewServer(cfg, http.NewRouter(db))

	// start http
	serv.Run()
}
