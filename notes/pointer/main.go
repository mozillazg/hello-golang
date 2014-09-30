package main

import "fmt"

// 函数接受一个指针
// *int 指针类型
func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func square(x *float64) {
	*x = *x * *x
}

func swap(x *int, y *int) {
	tmp := new(int)
	*tmp = *x
	*x = *y
	*y = *tmp
}

func main() {
	x := 5
	zero(&x) // &x 获取 x 的指针地址
	fmt.Println(x)

	y := new(int) // 指针
	one(y)
	fmt.Println(*y) // *指针变量 获取值

	z := 1.1
	square(&z)
	fmt.Println(z)

	a := 1
	b := 2
	fmt.Println("before: ", a, b)
	swap(&a, &b)
	fmt.Println("after: ", a, b)
}
