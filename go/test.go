package main

import (
	"fmt"
	"os"
	"time"
)

type tt struct {
	name string
	age int
}

func main() {
	dir, _ := os.Getwd()

	fmt.Println(dir)

}

func test(args ...interface{}) {
	for i:=0;i<=1000000;i++ {
		i = i + 1
	}
	time.Sleep(time.Second * 30)
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
