package config

import (
	"fmt"
	"log/slog"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	defaultServerPort         = 8080
	defaultJWTExpirationHours = 72
)

type (
	Config struct {
		HTTP HTTPConfig 
		DB   DBConfig
		JWT  JWTConfig
	}

	HTTPConfig struct {
		Port string `yaml:"port" env-default:"5432"`
	}
	DBConfig struct {
		Host     string `env:"DB_HOST" env-default:"localhost"`
		Username string `env:"DB_USERNAME" env-default:"localhost"`
		Password string `env:"DB_PASSWORD" env-default:"localhost"`
		Database string `env:"DB_DATABASE" env-default:"test_db"`
		Port     string `env:"DB_PORT" env-default:"5432"`
		SSLMode  string `yaml:"db.sslmode" env-default:"false"`
	}
	JWTConfig struct {
		SecretKey string `json:"secret_key"`
	}
)

func Init(configFile string) *Config {

	cfg := &Config{}

	err := cleanenv.ReadConfig(fmt.Sprintf("./config/%s.yml", configFile), cfg)
	if err != nil {
		slog.Error("error reading config file", "error", err)
	}

	return cfg
}
