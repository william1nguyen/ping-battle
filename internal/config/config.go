package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	RedisAddr  string
	SessionTTL int
}

var Cfg Config

func LoadConfig() {
	godotenv.Load()

	Cfg = Config{
		RedisAddr:  getEnv("REDIS_ADDR", "localhost:6379"),
		SessionTTL: coerceInt("SESSION_TTL", 360),
	}
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
