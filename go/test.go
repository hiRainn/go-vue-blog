package main

import (
	"fmt"
)

type tt struct {
	name string
	age int
}

func main() {
	test := []int{}
	fmt.Println(test == nil)
}

func test(args ...interface{}) {
	fmt.Println(len(args))
}

func reverseString(s []byte)  []byte {

	l := len(s)
	for i:=0; i< len(s) / 2; i ++ {
		tmp := s[l - i - 1];
		s[l - i - 1] = s[i]
		s[i] = tmp
	}
	return s
}
