package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Port    string `env:"PORT"`
	DBHost  string `env:"DB_HOST"`
	DBPort  string `env:"DB_PORT"`
	DBUser  string `env:"DB_USER"`
	DBPass  string `env:"DB_PASS"`
	DBName  string `env:"DB_NAME"`
	SSLMode bool   `env:"DB_SSL_MODE"`
}

func Parse() (Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
