package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func LoadConfig() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Warn().Msg("No .env file found, using system environment variables")
	}

	// Parse Redis variable
	if err := env.Parse(&RedisConfig); err != nil {
		log.Fatal().Err(err).Msg("Failed to parse environment variables for redis")
	}

	// Parse PostgreSQL variable
	if err := env.Parse(&PostgreSQLConfig); err != nil {
		log.Fatal().Err(err).Msg("Failed to parse environment variables for database")
	}

	// Parse GORM variable
	if err := env.Parse(&GormConfig); err != nil {
		log.Fatal().Err(err).Msg("Failed to parse environment variables for GORM")
	}
    
    log.Info().Msg("Configuration loaded successfully")
}