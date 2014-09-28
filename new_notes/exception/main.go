package main

import "fmt"

func main() {
	// 捕获异常
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	// 抛出异常
	panic("PANIC")
}
