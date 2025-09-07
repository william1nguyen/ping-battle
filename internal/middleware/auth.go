package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/william1nguyen/ping-battle/internal/redis"
)

func GameSessionMiddleware(c *gin.Context) {
	sessionID := c.Query("sessionID")

	if sessionID == "" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"error": "missing session ID",
			},
		)
	}

	key := fmt.Sprintf("session:%s", sessionID)
	username, err := redis.Rdb.Get(key).Result()

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"error": "invalid session"},
		)
		return
	}

	c.Set("username", username)
	c.Next()
}
