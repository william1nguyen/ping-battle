package redis

import (
	"github.com/go-redis/redis"
	"github.com/william1nguyen/ping-battle/internal/config"
)

var (
	Rdb *redis.Client
)

func Init() error {
	Rdb = redis.NewClient(&redis.Options{
		Addr: config.Cfg.RedisAddr,
	})

	_, err := Rdb.Ping().Result()
	return err
}
