package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServerAddress string `env:"RUN_ADDRESS" envDefault:":8080"`
	DatabaseDSN   string `env:"DATABASE_URI" envDefault:"postgres://user:pass@database/db"`
	MigrateFile   string `env:"MIGRATE" envDefault:"migrations"`
	Secret        string `env:"SECRET" envDefault:"secret"`
	LogLevel      string `env:"LOG_LEVEL" envDefault:"INFO"`
	LogColor      bool   `env:"LOG_COLOR" envDefault:"false"`
	ACCRUAL_URI   string `env:"ACCRUAL_SYSTEM_ADDRESS" envDefault:"accrual:8080"`
}

func NewFromEnv() (*Config, error) {
	var config Config
	err := env.Parse(&config)
	if err != nil {
		return &config, err
	}
	return &config, nil
}
