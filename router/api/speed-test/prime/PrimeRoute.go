package prime

import (
	"math"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func PrimeRoute(g *gin.RouterGroup) {
	g.GET("/prime", func(c *gin.Context) {
		incl := c.Query("include")
		to, err := strconv.Atoi(c.Query("to"))

		if err != nil {
			to = 10000000
		}

		start := time.Now()
		ps := SpeedTest(int32(to))
		elapsed := time.Since(start).String()

		res := gin.H {
			"numbersTested": gin.H {
				"start": 0,
				"to": to,
			},
			"elapsedTime": elapsed,
			"amount": len(ps),
		}

		if incl == "true" {res["results"] = ps}

		c.SecureJSON(200, res)
	})
}

func SpeedTest(to int32) []int32 {
	p := []int32{}

	for i := int32(0); i < to; i++ {
		if prime(i) {p = append(p, i)}
	}
	return p
}

func prime(n int32) bool {
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if int(n) % i == 0 {return false}
	}
	return true
}

