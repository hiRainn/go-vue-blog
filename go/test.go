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
	var res3 []int

	fmt.Println(res3 == nil)



}

func test(a map[int]string) {
	delete(a,1)
}


