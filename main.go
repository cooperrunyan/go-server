package main

import (
	"fmt"
	"http/router"
	"http/utils"
)

func main() {
	r := router.Init()
  e := r.Run("0.0.0.0:" + utils.GetPort())
	if e != nil {fmt.Println(e)}
}
