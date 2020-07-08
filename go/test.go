package main

import (
	"fmt"
	//"strings"
)

type tt struct {
	name string
	age int
}

func main() {
	data := []tt{
		{name: "heihei",age: 18},
		{name: "heihei",age: 18},
	}
	test(data)
}

func test(data interface{}) {
	fmt.Println(data)
}
