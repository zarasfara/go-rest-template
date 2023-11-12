package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/zarasfara/go-rest-template/internal/app"
	"github.com/zarasfara/go-rest-template/internal/config"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading .env file: %s", err.Error())
	}
	cfg := config.Init(os.Getenv("APP_ENV"))

	app.Run(cfg)
}