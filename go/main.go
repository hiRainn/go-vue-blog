// geektutu.com
// main.go
package main

import (
	"blog/model"
	"fmt"
)

func main() {
	//r := route.Initroute()
	//r.Run()
	 data := model.GetList()
	fmt.Println(data)

}