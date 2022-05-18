package middleware

import "github.com/gin-gonic/gin"

func Cors() func (c *gin.Context) {
	return func (c *gin.Context) {
		c.Writer.Header().Set("access-control-allow-credentials", "true")
		c.Writer.Header().Set("access-control-allow-headers", "*")
		c.Writer.Header().Set("access-control-allow-methods", "*")
		c.Writer.Header().Set("access-control-allow-origin", "*")
		c.Writer.Header().Set("cache-control", "no-cache")
		c.Writer.Header().Set("content-type", "application/json")

		c.Next()
	}
}
