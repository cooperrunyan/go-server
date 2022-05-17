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
	use(g, api.Ping)

	return g
}

func use(g *gin.Engine, h func(*gin.Engine)) { h(g) }
