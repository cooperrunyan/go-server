package api

import (
	"github.com/gin-gonic/gin"
)

func Ping(g *gin.Engine) {
	e := g.Group("/ping")

	e.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
