package api

import (
	"github.com/gin-gonic/gin"
)

func Ping(g *gin.Engine) {
	g.GET("/ping", func(c *gin.Context) {
		c.SecureJSON(200, gin.H{
			"message": "pong",
		})
	})
}
