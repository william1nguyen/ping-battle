package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisAddr  string
	SessionTTL int
}

var Cfg Config

func LoadConfig() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("failed to run server: %v", err)
	}

	Cfg = Config{
		RedisAddr:  getEnv("REDIS_ADDR", "localhost:6379"),
		SessionTTL: coerceInt("SESSION_TTL", 360),
	}

	return nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func coerceInt(key string, fallback int) int {
	if valueStr, ok := os.LookupEnv(key); ok {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}

	return fallback
}
