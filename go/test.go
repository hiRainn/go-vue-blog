package main

import "fmt"

type tt struct {
	name string
	age int
}

func main() {
	aaa := make([]interface{},4)

	aaa[0] = "10"
	aaa[1] = 100
	aaa[2] = 20
	aaa[3] = 5
	test(&aaa)

	fmt.Println(aaa)

}

func test(arr *[]interface{}) {
	*arr = append((*arr)[:1],(*arr)[2:]...)
	fmt.Println(arr)
}


