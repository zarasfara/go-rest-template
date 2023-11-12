package config

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
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
		Port string
	}
	DBConfig struct {
		Host     string
		Username string
		Password string
		Database string
		Port     string
		SSLMode  string
	}
	JWTConfig struct {
		SecretKey string `json:"secret_key"`
	}
)

func Init(configFile string) *Config {
	populateDefaults()

	if err := parseConfig(configFile); err != nil {
		logrus.Fatalf("cannot parse config: %s", err.Error())
	}

	var cfg *Config

	setFromEnv(cfg)

	return cfg
}

// unmarshal set parametres from config.yml file.
func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		logrus.Fatalf("error read config: %s", err)
	}

	return nil
}

// setFromEnv set parametres from .env.
func setFromEnv(cfg *Config) {
	cfg.JWT.SecretKey = os.Getenv("JWT_SIGNING_TOKEN")
}

func populateDefaults() {
	viper.SetDefault("http.port", defaultServerPort)
}

// parseConfig parse config from file config.yml.
func parseConfig(configFile string) error {
	viper.AddConfigPath("configs")
	viper.SetConfigName(configFile)

	return viper.ReadInConfig()
}
