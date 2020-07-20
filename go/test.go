package main

import (
	"fmt"
	"strconv"
)

type tt struct {
	name string
	age int
}

func main() {
	s := make(map[string]interface{})
	s["sort"] = ""

	sort,_ :=strconv.Atoi(s["sort"].(string))
	fmt.Println(sort)
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
