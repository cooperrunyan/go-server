package utils

import (
	"os"
	"strings"
)

func GetPort() string {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0] == "PORT" { return pair[1] }
	}
	return "3000"
}
