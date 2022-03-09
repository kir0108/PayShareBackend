package main

import (
	"time"

	"github.com/caarlos0/env"
)

type config struct {
	Production           bool          `env:"PRODUCTION" envDefault:"false"`
	Port                 string        `env:"PORT" envDefault:"80"`
	PostgresUrl          string        `env:"POSTGRES_URL" envDefault:"postgres://postgres:carnetpass@db:5432/postgres?sslmode=disable"`
	RedisUrl             string        `env:"REDIS_URL" envDefault:"redis:6379"`
	JwtTTL               time.Duration `env:"TOKEN_TTL" envDefault:"20m"`
	Secret               string        `env:"SECRET" envDefault:"test_secret"`
	SessionTTl           time.Duration `env:"SESSION_TTL" envDefault:"168h"`
	SessionCleanupPeriod time.Duration `env:"SESSION_CLEANUP_PERIOD" envDefault:"60s"`
	SessionWindowPeriod  time.Duration `env:"SESSION_WINDOW_PERIOD" envDefault:"60s"`
	SessionTokenLength   int           `env:"SESSION_TOKEN_LENGTH" envDefault:"32"`
	CodeHideTTL          time.Duration `env:"CodeHideTTL" envDefault:"60m"`
	CodeExpiredTTL       time.Duration `env:"CodeExpiredTTL" envDefault:"70m"`
	CodeLength           int           `env:"CODE_LENGTH" envDefault:"5"`
	SecretPhrase         string        `env:"SECRET_PHRASE"`
	IsTest               bool          `env:"IS_TEST" envDefault:"false"`
}

func getConfig() (*config, error) {
	c := &config{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}

	return c, nil
}
