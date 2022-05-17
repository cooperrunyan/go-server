package router

import (
	"http/router/api"
	"http/router/api/code"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	g := gin.Default()

	use(g, api.Echo)
	use(g, code.Router)
	use(g, api.Proxy)

	g.GET("/", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "Running",
		})
	})

	return g
}

func use(g *gin.Engine, h func(*gin.Engine)) { h(g) }
