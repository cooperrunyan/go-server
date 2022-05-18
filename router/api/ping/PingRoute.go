package ping

import (
	"github.com/gin-gonic/gin"
)

func PingRoute(g *gin.Engine) {
	g.GET("/ping", func(c *gin.Context) {
		c.SecureJSON(200, gin.H{
			"message": "pong",
		})
	})
}
