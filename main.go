package main

import (
	"fmt"
	"http/router"
	"http/utils"
)

func main() {
	r := router.Init()

  e := r.Run(":" + utils.GetPort())
	if e != nil {fmt.Println(e)}
}
