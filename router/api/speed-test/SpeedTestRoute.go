package speedtest

import (
	"http/router/api/speed-test/iteration"
	"http/router/api/speed-test/prime"
	"http/utils"

	"github.com/gin-gonic/gin"
)

func SpeedTestRoute(r *gin.Engine) {
	g := r.Group("/speed-test")

	utils.AttachRouteToGroup(g, prime.PrimeRoute)
	utils.AttachRouteToGroup(g, iteration.IterationRoute)
}
