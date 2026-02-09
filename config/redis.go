package config

import (
	"context"
	"crypto/tls"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type redisConfig struct {
	Host      string `env:"REDIS_HOST"`
	Port      int    `env:"REDIS_PORT"`
	Username  string `env:"REDIS_USER"`
	Password  string `env:"REDIS_PASS"`
	Database  int    `env:"REDIS_DB_INDEX"`
	EnableTls bool   `env:"REDIS_USE_SSL"`
}

var RedisConfig redisConfig

var redisConnection *redis.Client
var redisOnce sync.Once

func GetRedisConnection() *redis.Client {
	redisOnce.Do(func() {
		redisOpts := &redis.Options{
			Addr:     fmt.Sprintf("%s:%d", RedisConfig.Host, RedisConfig.Port),
			Password: RedisConfig.Password,
			DB:       RedisConfig.Database,
		}
		if RedisConfig.EnableTls {
			redisOpts.TLSConfig = &tls.Config{
				InsecureSkipVerify: true,
			}
		}
		client := redis.NewClient(redisOpts)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		result, err := client.Ping(ctx).Result()
		if err != nil {
			log.Fatal().Err(err).Msg("Error ping redis")
		}
		log.Debug().Any("Status", result).Msg("Redis status")
		redisConnection = client
	})

	return redisConnection
}
