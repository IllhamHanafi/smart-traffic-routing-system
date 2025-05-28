package config

import "time"

type Database struct {
	Library      string `env:"LIBRARY" envDefault:"sqlc"`
	Host         string `env:"HOST" envDefault:"localhost"`
	Port         int    `env:"PORT" envDefault:"5432"`
	Username     string `env:"USERNAME"`
	Password     string `env:"PASSWORD"`
	DatabaseName string `env:"DATABASE_NAME"`
}

type JWT struct {
	RSAPrivateKeyPath string        `env:"RSA_PRIVATE_KEY_PATH"`
	RSAPublicKeyPath  string        `env:"RSA_PUBLIC_KEY_PATH"`
	ExpiredIn         time.Duration `env:"EXPIRED_IN" envDefault:"15m"`
}
