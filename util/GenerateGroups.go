package util

import (
	"fmt"
	"http/types"
	"strings"

	"github.com/gin-gonic/gin"
)

func GenerateGroups() (gin.H, map[string]map[int]types.StatusCode) {
	gs := map[string]map[int]types.StatusCode {
		"100": {},
		"200": {},
		"300": {},
		"400": {},
		"500": {},
	}

	for c, d := range GetStatuses() {
		s := strings.Split(fmt.Sprint(c), "")[0] + "00"

		gs[s][c] = d
	}

	var d gin.H = gin.H{}

	for sl, g := range gs {
		d[sl] = gin.H {
			"codes": ParseStatuses(g),
		}
	}

	return d, gs
}
