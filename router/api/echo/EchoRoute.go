package echo

import (
	"io"

	"github.com/gin-gonic/gin"
)

func EchoRoute(g *gin.Engine) {
	g.Any("/echo", func(c *gin.Context) {
		a := map[string]string{}
		for k, v := range c.Request.URL.Query() { a[k] = v[0] }

		h := map[string]string{}
		for k, v := range c.Request.Header { h[k] = v[0] }

		var d []byte

		if c.Request.ContentLength > 0 {
			j, err := io.ReadAll(c.Request.Body)
			d = j

			if err != nil {
				c.SecureJSON(500, gin.H {
					"status": "error",
					"message": "Error parsing request body",
				})
				return
			}
		}

		rIp, trusted := c.RemoteIP()


		r := gin.H {
			"method": c.Request.Method,
			"data": string(d),
			"args": a,
			"headers": h,
			"path": c.Request.URL.Path,
			"requestContentLength": c.Request.ContentLength,
			"protocol": c.Request.Proto,
			"clientIp": c.ClientIP(),
			"remoteIp": rIp,
			"remoteAddress": c.Request.RemoteAddr,
			"isTrustedProxy": trusted,
			"requestURI": c.Request.RequestURI,
		}

		c.SecureJSON(200, r)
	})
}
