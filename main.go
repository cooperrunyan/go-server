package main

import (
	"http/router"
	"os"
	"strings"
)

func main() {
	r := router.Init()
	p := "3000"
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0] == "PORT" {
			p = pair[1]
		}
	}

	r.Run("0.0.0.0:" + p)
}
