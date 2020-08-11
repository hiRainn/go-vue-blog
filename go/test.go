package main

import "fmt"



type tt struct {
	name string
	age int
}

type Test int

func (t Test) Write() int{
	return 1
}

func (t *tt) test() {

}

func main() {
	var a Test

	fmt.Println(a.Write())



}

func test(a map[int]string) {
	delete(a,1)
}


