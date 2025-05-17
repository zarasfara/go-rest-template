package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/zarasfara/go-rest-template/internal/app"
	"github.com/zarasfara/go-rest-template/internal/config"
	"github.com/zarasfara/go-rest-template/internal/logger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error("error loading .env file", "error", err)
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	// init logger
	logger.InitLogger(env)

	cfg := config.Init(env)

	app.Run(cfg)
}