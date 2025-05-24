package config

import (
	"sync"

	"github.com/IllhamHanafi/smart-traffic-routing-system/shared-libs/config"
	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

var (
	once sync.Once
	cfg  Config
)

type Config struct {
	Environment string          `env:"ENVIRONMENT" envDefault:"development"`
	Port        int             `env:"PORT" envDefault:"8080"`
	Database    config.Database `envPrefix:"DB_"`
}

type Database struct {
	Library      string `env:"LIBRARY" envDefault:"sqlc"`
	Host         string `env:"HOST" envDefault:"localhost"`
	Port         int    `env:"PORT" envDefault:"5432"`
	Username     string `env:"USERNAME"`
	Password     string `env:"PASSWORD"`
	DatabaseName string `env:"DATABASE_NAME" envDefault:"api-gateway"`
}

func GetConfig() Config {
	once.Do(func() {
		err := env.Parse(&cfg)
		if err != nil {
			panic(err)
		}
	})
	return cfg
}
