package main

import (
	"fmt"
	"time"

	//"strings"
	"github.com/patrickmn/go-cache"
)

type tt struct {
	name string
	age int
}

func main() {
	c1 := cache.New(72*time.Hour,10*time.Minute)
	c1.Set("name","heihei",cache.DefaultExpiration)

	c2 := cache.New(72*time.Hour,10*time.Minute)
	v,_ :=c2.Get("name")
	v1,_ := c1.Get("name")
	println(v.(string),v1.(string))
}

func test(data interface{}) {
	fmt.Println(data)
}
