package api

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func Proxy(g *gin.Engine) {
	e := g.Group("/proxy")

	e.Any("/", func(c *gin.Context) {
		u := c.Request.URL.Query()["u"][0]

		if u == "" {
			c.JSON(400, gin.H {
				"message": "Send the url endpoint as a query param under the value: 'u'",
			})
			return
		}

		remote, err := url.Parse(u)
		if err != nil {
				c.JSON(400, gin.H {
				"message": "Not a valid url",
				"error": err,
			})
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			req.URL.Path = remote.Path
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	})
}
