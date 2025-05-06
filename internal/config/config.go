package config

import (
	"fmt"

	"github.com/caarlos0/env/v9"
	"github.com/joho/godotenv"
)

type Config struct {
	GRPCPort   string `env:"GRPC_PORT" envDefault:"50052"`
	DBHost     string `env:"DB_HOST,required"`
	DBPort     string `env:"DB_PORT" envDefault:"5432"`
	DBUser     string `env:"DB_USER,required" envDefault:"user"`
	DBPassword string `env:"DB_PASSWORD,required" envDefault:"1469"`
	DBName     string `env:"DB_NAME,required"`
	DBSSLMode  string `env:"DB_SSL_MODE" envDefault:"disable"`
	LogLevel   string `env:"LOG_LEVEL" envDefault:"info"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Warning: could not load .env file: %v. Falling back to environment variables.", err)
	}

	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
