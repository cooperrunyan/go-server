package router

import (
	"http/middleware"
	"http/router/api"
	"http/router/api/code"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	useRoute(r, api.Proxy)

	r.Use(middleware.Cors())

	useRoute(r, api.Echo)
	useRoute(r, code.Router)
	useRoute(r, api.Ping)

	return r
}

func useRoute(r *gin.Engine, h func(*gin.Engine)) { h(r) }
