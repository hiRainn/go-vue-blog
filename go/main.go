// geektutu.com
// main.go
package main

import (
	"blog/route"
	"github.com/gin-gonic/gin"
)

func main() {
	car :=route.Initroute()
	car.Use(gin.Logger())
	car.Run()
}



