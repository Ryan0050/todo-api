package config

import (
    "github.com/joho/godotenv"
    "github.com/caarlos0/env/v10"
    "github.com/rs/zerolog/log"
)

type dbConfig struct {
    Host     string `env:"DB_HOST"`
    Port     string `env:"DB_PORT"`
    User     string `env:"DB_USER"`
    Pass     string `env:"DB_PASS"`
    Name     string `env:"DB_NAME"`
}

var DBConfig dbConfig 

func LoadConfig() {
    if err := godotenv.Load(); err != nil {
        log.Warn().Msg("No .env file found, using system environment variables")
    }

    if err := env.Parse(&RedisConfig); err != nil {
        log.Fatal().Err(err).Msg("Failed to parse environment variables for redis")
    }

    if err := env.Parse(&DBConfig); err != nil {
        log.Fatal().Err(err).Msg("Failed to parse environment variables for database")
    }
}