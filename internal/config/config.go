package config

import (
	"github.com/caarlos0/env/v9"
)

type Config struct {
	GRPCPort   string `env:"GRPC_PORT" envDefault:"50051"`
	DBHost     string `env:"DB_HOST,required"`
	DBPort     string `env:"DB_PORT" envDefault:"5432"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBName     string `env:"DB_NAME,required"`
	DBSSLMode  string `env:"DB_SSL_MODE" envDefault:"disable"`
	LogLevel   string `env:"LOG_LEVEL" envDefault:"info"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
