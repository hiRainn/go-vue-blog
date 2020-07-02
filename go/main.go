// geektutu.com
// main.go
package main

import (
	"blog/model"
	"blog/route"
)

func main() {
	car :=route.Initroute()
	car.Run()
}

func init() {
	model.InitDb()
}

