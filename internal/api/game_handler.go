package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/william1nguyen/ping-battle/internal/service"
)

func SessionHanlder(c *gin.Context) {
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
