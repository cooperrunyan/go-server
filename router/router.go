package router

import (
	"http/middleware"
	"http/router/api/code"
	"http/router/api/echo"
	"http/router/api/ping"
	"http/router/api/proxy"
	speedtest "http/router/api/speed-test"
	"http/utils"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	utils.AttachRoute(r, proxy.ProxyRoute)

	r.Use(middleware.Cors())

	utils.AttachRoute(r, echo.EchoRoute)
	utils.AttachRoute(r, code.CodeRoute)
	utils.AttachRoute(r, ping.PingRoute)
	utils.AttachRoute(r, speedtest.SpeedTestRoute)

	return r
}


