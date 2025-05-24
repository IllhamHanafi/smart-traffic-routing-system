package config

type Database struct {
	Library      string `env:"LIBRARY" envDefault:"sqlc"`
	Host         string `env:"HOST" envDefault:"localhost"`
	Port         int    `env:"PORT" envDefault:"5432"`
	Username     string `env:"USERNAME"`
	Password     string `env:"PASSWORD"`
	DatabaseName string `env:"DATABASE_NAME" envDefault:"route-engine"`
}
