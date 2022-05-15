package util

import (
	"fmt"
	"http/types"

	"github.com/gin-gonic/gin"
)

func ParseStatuses(statuses map[int]types.StatusCode) gin.H {
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
