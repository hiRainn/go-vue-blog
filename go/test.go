package main

import (
	"fmt"
	"runtime"
	"time"
)

type tt struct {
	name string
	age int
}

func main() {
	for i:=0;i<2;i++ {
		tmp := int64(129)
		test()
		fmt.Printf("%p \r\n",&tmp)
	}

	defer func() {
		// 发生宕机时，获取panic传递的上下文并打印
		err := recover()
		switch err.(type) {
		case runtime.Error: // 运行时错误
			fmt.Println("runtime error:", err)
		default: // 非运行时错误
			fmt.Println("error:", err)
		}
	}()


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
