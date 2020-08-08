package main

import "fmt"

type tt struct {
	name string
	age int
}

func (t *tt) test() {
	fmt.Println("aaa")
}

func main() {
	 a := []int{1,2,3}
	 test(a)

	 fmt.Println(a)

}

func test(a []int) {
	a[0] = 10
}


