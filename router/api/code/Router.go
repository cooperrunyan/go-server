package code

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Router(g *gin.Engine) {
	c := g.Group("/code")

	c.GET("/", func(c *gin.Context) {
		c.JSON(200, parseStatuses(getStatuses()))
	})

	c.GET("/list", func(c *gin.Context) {
	  d := map[int]string{}
		for c, data := range getStatuses() {
			d[c] = data.Phrase
		}

		c.JSON(200, d)
	})

	c.GET("/group", func(c *gin.Context) {
		d, _ := generateGroups();
		c.JSON(200, d)
	})

	c.GET("/group/:group", func(c *gin.Context) {
		g := strings.Split(c.Param("group"), "")[0]
		_, gs := generateGroups();

		num, err := strconv.Atoi(g)

		if err != nil {
			c.JSON(400, gin.H {
				"error": "Given code could not be parsed to an integer.",
			})
			return
		}

		if num > 5 || num < 1 {
			c.JSON(400, gin.H {
				"error": "Code should be within 100 and 500 range",
			})
			return
		}

		c.JSON(200, gs[g + "00"])
	})

	c.GET("/code/:code", func(c *gin.Context) {
		q := c.Param("code")
		code, err := strconv.Atoi(q)

		if err != nil {
			c.JSON(400, gin.H {
				"ok": false,
				"error": "Given code: '" + q + "' could not be parsed to an integer.",
			})
			return
		}

		d := getStatuses()[code]

		if d.Phrase == "" {
			c.JSON(400, gin.H {
				"ok": false,
				"message": "Status code not found",
			})
			return
		}

		c.JSON(200, gin.H {
			"ok": true,
			"data": gin.H {
				"code": d.Code,
				"phrase": d.Phrase,
				"description": d.Description,
				"spec": d.Spec,
				"specLink": d.SpecLink,
			},
		})
	})

	c.GET("/respond/:code", func(c *gin.Context) {
		q := c.Param("code")
		code, err := strconv.Atoi(q)

		if err != nil {
			c.JSON(400, gin.H {
				"ok": false,
				"error": "Given code: '" + q + "' could not be parsed to an integer.",
			})
			return
		}

		d := getStatuses()[code]

		if d.Phrase == "" {
			c.JSON(400, gin.H {
				"ok": false,
				"message": "Status code not found",
			})
			return
		}

		c.JSON(code, gin.H {
			"ok": true,
			"code": code,
			"message": getStatuses()[code].Description,
		})
	})
}
