package code

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func parseStatuses(statuses map[int]StatusCode) gin.H {
	var d gin.H = gin.H{}

	for code, data := range statuses {
		d[fmt.Sprint(code)] = gin.H {
			"code": data.Code,
			"phrase": data.Phrase,
			"description": data.Description,
			"spec": data.Spec,
			"specLink": data.SpecLink,
		}
	}

	return d
}
