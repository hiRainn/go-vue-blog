package main

import (
	"fmt"
	"time"
)
type Link struct{
	ele interface{}
	link *Link
}

func main() {
	star := time.Now()
	fmt.Println("hello world")
	end := time.Since(star)
	fmt.Println(end)

}