package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:":8080"`
	DatabaseDSN   string `env:"DB_DSN" envDefault:"postgres://user:pass@database/db"`
	MigrateFile   string `env:"MIGRATE" envDefault:"migrations"`
	Secret        string `env:"SECRET" envDefault:"secret"`
	LogLevel      string `env:"LOG_LEVEL" envDefault:"INFO"`
	LogColor      bool   `env:"LOG_COLOR" envDefault:"false"`
}

func NewFromEnv() (*Config, error) {
	var config Config
	err := env.Parse(&config)
	if err != nil {
		return &config, err
	}
	return &config, nil
}
