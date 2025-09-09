package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/william1nguyen/ping-battle/internal/api"
	"github.com/william1nguyen/ping-battle/internal/config"
	"github.com/william1nguyen/ping-battle/internal/redis"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		panic(err.Error())
	}

	if err := redis.Init(); err != nil {
		log.Fatalf("failed to init redis: %v", err)
	}

	r := gin.Default()
	api.RegisterGameRoutes(r)

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
