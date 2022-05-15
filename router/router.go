package router

import (
	"http/router/api"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	g := gin.Default()

	use(g, api.Echo)
	use(g, api.Code)
	use(g, api.Proxy)

	return g
}

func use(g *gin.Engine, h func(*gin.Engine)) { h(g) }
