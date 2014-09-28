package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func f() (int, int) {
	return 5, 6
}

// 不定参数
func f2(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

// 闭包
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// 递归
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}

func a() {
	fmt.Println("a")
}

func b() {
	fmt.Println("b")
}

// 延迟, 在程序 return 前执行 defer 后的表达式
func aDefer() {
	defer b() // 在程序 return 前执行 defer 后的表达式
	a()
}

func main() {
	// n := []float64{10, 20, 31.2, 40.0, 1.2}
	// fmt.Println(average(n))

	// f()
	// x, y := f()
	// fmt.Println(x, y)

	// fmt.Println(f2(1, 2, 3, 4, 5, 6))
	// arr := []int{1, 2, 3, 4, 5, 6}
	// fmt.Println(f2(arr...))

	// nextEven := makeEvenGenerator()
	// fmt.Println(nextEven())
	// fmt.Println(nextEven())
	// fmt.Println(nextEven())
	// fmt.Println(nextEven())
	// fmt.Println(factorial(5))
	aDefer()
}
