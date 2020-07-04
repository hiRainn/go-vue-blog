package main

import "fmt"

func main() {
	test := func () (res int) {
		defer func() {
			res += 1
		}()
		return 20
	}()

	fmt.Println(test)
}