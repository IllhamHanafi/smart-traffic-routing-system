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
	Environment         string          `env:"ENVIRONMENT" envDefault:"development"`
	Port                int             `env:"PORT" envDefault:"8080"`
	Database            config.Database `envPrefix:"DB_"`
	BcryptPasswordRound int             `env:"BCRYPT_PASSWORD_ROUND" envDefault:"12"`
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
