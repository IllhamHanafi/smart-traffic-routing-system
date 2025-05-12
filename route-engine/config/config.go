package config

import (
	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Environment string   `env:"ENVIRONMENT" envDefault:"development"`
	Port        int      `env:"PORT" envDefault:"8080"`
	Database    Database `envPrefix:"DB_"`
}

type Database struct {
	Host     string `env:"HOST" envDefault:"localhost"`
	Port     int    `env:"PORT" envDefault:"5432"`
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	Name     string `env:"NAME" envDefault:"route-engine"`
}

func ParseConfig() *Config {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}
	return &cfg
}
