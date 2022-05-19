package iteration

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func IterationRoute(g *gin.RouterGroup) {
	g.GET("/iteration", func(c *gin.Context) {
		t, er := strconv.Atoi(c.Query("i"))
		iter := t

		if er != nil {iter = 1000000}

		s := time.Now()
		ctr := 0

		for i := 0; i <= iter; i++ {
			ctr++
		}

		e := time.Since(s).String()

		c.SecureJSON(200, gin.H {
			"numbersTested": gin.H {
				"start": 0,
				"to": t,
			},
			"elapsedTime": e,
			"iterations": iter,
		})
	})
}
