package main

import (
	"fmt"
	"strings"
)

type tt struct {
	name string
	age int
}


func main() {
	var a = "1,2,4,5"

	arr := strings.Split(a,",")

	for _,v := range  arr {
		fmt.Println(v)
	}



}

func test(a map[int]string) {
	delete(a,1)
}


