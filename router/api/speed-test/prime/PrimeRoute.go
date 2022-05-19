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
		ps := SieveOfEratosthenes(to)
		elapsed := time.Since(start).String()

		res := gin.H {
			"data": gin.H {
				"tested": to,
				"amount": len(ps),
				"elapsedTime": elapsed,
			},
		}

		if incl == "true" {res["results"] = ps}

		c.SecureJSON(200, res)
	})
}

func SieveOfEratosthenes(to int) []int {
	f := make([]bool, to)
	ps := []int{}
	for i := 2; i <= int(math.Sqrt(float64(to))); i++ {
		if !f[i] {
			for j := i * i; j < to; j += i { f[j] = true }
		}
	}
	for i := 2; i < to; i++ { if !f[i] { ps = append(ps, i) } }
	return ps
}
