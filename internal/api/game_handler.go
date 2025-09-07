package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/william1nguyen/ping-battle/internal/middleware"
	"github.com/william1nguyen/ping-battle/internal/service"
)

func RegisterGameRoutes(r *gin.Engine) {
	game := r.Group("/game")
	{
		game.POST("/session", getSessionHandler)
		game.GET("/ping", middleware.GameSessionMiddleware, pingHandler)
		game.GET("/top", getTopUsersHanlder)
		game.GET("/count", getCountUniqueHandler)
	}
}

func getSessionHandler(c *gin.Context) {
	username := c.PostForm("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "username required",
		})
		return
	}
	user := service.CreateSession(username)
	c.JSON(http.StatusOK, user)
}

func pingHandler(c *gin.Context) {
	username := c.GetString("username")

	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "username required",
		})
		return
	}

	msg, err := service.ProcessPing(username)

	if err != nil {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func getTopUsersHanlder(c *gin.Context) {
	top, err := service.GetTopUsers()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, top)
}

func getCountUniqueHandler(c *gin.Context) {
	count, err := service.GetUniqueCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"unique_users": count,
	})
}
