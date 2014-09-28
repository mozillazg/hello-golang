package main

import "fmt"

func sum(args ...int) int {
	total := 0
	for _, n := range args {
		total += n
	}
	return total
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

func max(args ...int) int {
	m := args[0]
	for _, v := range args {
		if v > m {
			m = v
		}
	}
	return m
}

func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (m uint) {
		m = i
		i += 2
		return
	}
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(isEven(1))
	fmt.Println(isEven(2))
	fmt.Println(max(1, 5, 0, 7, 3))
	odd := makeOddGenerator()
	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println("")
	fmt.Println(fib(5))
}
