package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/william1nguyen/ping-battle/internal/api"
	"github.com/william1nguyen/ping-battle/internal/config"
	"github.com/william1nguyen/ping-battle/internal/redis"
)

func main() {
	config.LoadConfig()

	if err := redis.Init(); err != nil {
		log.Fatalf("failed to init redis: %v", err)
	}

	r := gin.Default()
	api.RegisterGameRoutes(r)

	r.Run()
}
