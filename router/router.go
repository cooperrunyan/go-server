package router

import (
	"http/middleware"
	"http/router/api/code"
	"http/router/api/echo"
	"http/router/api/ping"
	"http/router/api/proxy"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	useRoute(r, proxy.ProxyRoute)

	r.Use(middleware.Cors())

	useRoute(r, echo.EchoRoute)
	useRoute(r, code.CodeRoute)
	useRoute(r, ping.PingRoute)

	return r
}

func useRoute(r *gin.Engine, h func(*gin.Engine)) { h(r) }
