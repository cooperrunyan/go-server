package api

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

const urlParamName = "u"

func Proxy(g *gin.Engine) {
	g.Any("/proxy", func(c *gin.Context) {
		u := c.Request.URL.Query()[urlParamName][0]

		if u == "" {
			c.SecureJSON(400, gin.H {
				"message": `Send the url endpoint as a query param under the value: '{{urlParamName}}'`,
			})
			return
		}

		remote, err := url.Parse(u)
		if err != nil {
				c.SecureJSON(400, gin.H {
				"message": "Not a valid url",
				"error": err,
			})
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.Method = c.Request.Method
			req.URL = remote

			if c.Request.Method != "GET" {req.Body = c.Request.Body}
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	})
}
