package code

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func generateGroups() (gin.H, map[string]map[int]StatusCode) {
	gs := map[string]map[int]StatusCode {
		"100": {},
		"200": {},
		"300": {},
		"400": {},
		"500": {},
	}

	for c, d := range getStatuses() {
		s := strings.Split(fmt.Sprint(c), "")[0] + "00"

		gs[s][c] = d
	}

	var d gin.H = gin.H{}

	for sl, g := range gs {
		d[sl] = gin.H {
			"codes": parseStatuses(g),
		}
	}

	return d, gs
}
