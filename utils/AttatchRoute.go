package utils

import "github.com/gin-gonic/gin"

func AttachRoute(r *gin.Engine, h func(*gin.Engine)) { h(r) }
func AttachRouteToGroup(r *gin.RouterGroup, h func(*gin.RouterGroup)) { h(r) }
