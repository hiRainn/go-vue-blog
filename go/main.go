// geektutu.com
// main.go
package main

import (
	"blog/config"
	"fmt"
)



func main() {
	//r := route.Initroute()
	//r.Run()
	conf := config.GetConf()
	fmt.Println(conf.DB_username)

}